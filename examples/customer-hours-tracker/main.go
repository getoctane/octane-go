package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/getoctane/octane-go"
	"github.com/gin-gonic/gin"
)

var (
	// Instantiate the Octane API client
	client = octane.NewClient(os.Getenv("OCTANE_API_KEY"))

	// Application settings and defaults
	port          = envOrDefaultInt("APP_PORT", 3000)
	bind          = envOrDefaultStr("APP_BIND", "127.0.0.1")
	meterName     = envOrDefaultStr("OCTANE_METER_NAME", "hours")
	pricePlanName = envOrDefaultStr("OCTANE_PRICE_PLAN_NAME", "standard")
	pricePlanRate = envOrDefaultInt("OCTANE_PRICE_PLAN_RATE", 30)
)

func buildRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFile("/", "./public/index.html")
	r.StaticFile("/index.html", "./public/index.html")
	r.StaticFile("/scripts.js", "./public/scripts.js")
	r.StaticFile("/styles.css", "./public/styles.css")
	r.GET("/api/customers", handleGetAPICustomers)
	r.POST("/api/customers", handlerPostAPICustomers)
	r.DELETE("/api/customers/:name", handlerDeleteAPICustomerName)
	r.POST("/api/hours", handlerPostAPIHours)
	return r
}

func handleGetAPICustomers(c *gin.Context) {
	fmt.Println("[octane] Listing customers in account")
	if customers, _, err := client.Customers.List(); err != nil {
		fmt.Println("[octane] Error listing customers in account")
		badRequest(c, err)
	} else {
		if customers == nil {
			customers = []octane.Customer{}
		}
		fmt.Printf("[octane] Total customers: %d\n", len(customers))
		c.JSON(200, customers)
	}
}

func handlerPostAPICustomers(c *gin.Context) {
	var data struct{ Name string }
	if err := c.BindJSON(&data); err != nil {
		badRequest(c, err)
		return
	}
	fmt.Printf("[octane] Attempting to create new customer \"%s\"\n", data.Name)
	createCustomerArgs := octane.CreateCustomerArgs{
		Name: data.Name,
		MeasurementMappings: []octane.CustomerMeasurementMappingInputArgs{
			{
				Label:      "customer_name",
				ValueRegex: data.Name,
			},
		},
	}
	if _, _, err := client.Customers.Create(createCustomerArgs); err != nil {
		fmt.Printf("[octane] Error creating customer \"%s\"\n", data.Name)
		badRequest(c, err)
		return
	}
	fmt.Printf("[octane] Customer \"%s\" successfully created\n", data.Name)
	fmt.Printf("[[octane] Attempting to subscribe customer \"%s\" to price plan \"%s\"\n",
		data.Name, pricePlanName)
	createSubscriptionArgs := octane.CreateSubscriptionArgs{
		PricePlanName: pricePlanName,
	}
	if _, _, err := client.Customers.CreateSubscription(data.Name, createSubscriptionArgs); err != nil {
		fmt.Printf(
			"[octane] Error subscribing customer \"%s\" to price plan \"%s\"\n",
			data.Name, pricePlanName)
		badRequest(c, err)
		return
	}
	fmt.Printf(
		"[octane] Successfully subscribed customer \"%s\" to price plan \"%s\"\n",
		data.Name, pricePlanName)
	success(c, 201)
}

func handlerDeleteAPICustomerName(c *gin.Context) {
	name := c.Param("name")
	fmt.Printf(
		"[octane] Attempting to unsubscribe customer \"%s\" from price plan \"%s\"\n",
		name, pricePlanName)
	if _, err := client.Customers.DeleteSubscription(name, octane.DeleteSubscriptionArgs{}); err != nil {
		fmt.Printf(
			"[octane] Error unsubscribing customer \"%s\" from price plan \"%s\"\n",
			name, pricePlanName)
		badRequest(c, err)
		return
	}
	fmt.Printf(
		"[octane] Customer \"%s\" successfully unsubscribed from price plan \"%s\"\n",
		name, pricePlanName)
	if _, err := client.Customers.Delete(name); err != nil {
		fmt.Printf("[octane] Error deleting customer \"%s\"\n", name)
		badRequest(c, err)
		return
	}
	fmt.Printf("[octane] Customer \"%s\" successfully deleted\n", name)
	success(c, 200)
}

func handlerPostAPIHours(c *gin.Context) {
	var data struct {
		Name  string
		Hours interface{}
	}
	if err := c.BindJSON(&data); err != nil {
		badRequest(c, err)
		return
	}
	var value float64
	if v, err := strconv.ParseFloat(fmt.Sprintf("%v", data.Hours), 8); err != nil {
		badRequest(c, err)
		return
	} else {
		value = v
	}
	measurement := octane.Measurement{
		Value: value,
		Labels: map[string]string{
			"customer_name": data.Name,
		},
		MeterName: meterName,
	}
	fmt.Printf("[octane] Attempting to create measurement for customer \"%s\"\n", data.Name)
	if _, _, err := client.Measurements.Create(measurement); err != nil {
		fmt.Printf("[octane] Error creating measurement for customer \"%s\"\n", data.Name)
		badRequest(c, err)
		return
	}
	fmt.Printf("[octane] Measurement for customer \"%s\" successfully created\n", data.Name)
	success(c, 201)
}

func success(c *gin.Context, status int) {
	c.JSON(status, gin.H{
		"code":    status,
		"message": "success",
	})
}

func badRequest(c *gin.Context, err error) {
	fmt.Printf("[server] Error: %s\n", err.Error())
	c.JSON(400, gin.H{
		"code":    400,
		"status":  err.Error(),
		"message": err.Error(),
	})
}

func checkOctaneAPIKey() {
	fmt.Println("[octane] Checking if OCTANE_API_KEY is present in environment")
	if key := os.Getenv("OCTANE_API_KEY"); key == "" {
		crash("Must set OCTANE_API_KEY.")
	}
	fmt.Println("[octane] OCTANE_API_KEY is present in environment")
}

func checkOctaneResources() {
	fmt.Println("[octane] Ensuring all required resources exist")
	checkOctaneResourceMeter()
	checkOctaneResourcePricePlan()
}

func checkOctaneResourceMeter() {
	fmt.Printf("[octane] Checking if meter \"%s\" exists\n", meterName)
	if _, _, err := client.Meters.Retrieve(meterName); err == nil {
		fmt.Printf("[octane] Meter \"%s\" already exists\n", meterName)
		return
	}
	fmt.Printf("[octane] Meter \"%s\" does not exist, creating\n", meterName)
	meterInputArgs := octane.MeterInputArgs{
		Name:           meterName,
		DisplayName:    "Number of hours worked",
		MeterType:      "COUNTER",
		UnitName:       "hour",
		IsIncremental:  true,
		ExpectedLabels: []string{"customer_name"},
	}
	if _, _, err := client.Meters.Create(meterInputArgs); err != nil {
		crash(fmt.Sprintf("Unable to create meter: %s", err.Error()))
	}
	fmt.Printf("[octane] Meter \"%s\" successfully created\n", meterName)
}

func checkOctaneResourcePricePlan() {
	fmt.Printf("[octane] Checking if price plan \"%s\" exists\n", pricePlanName)
	if _, _, err := client.PricePlans.Retrieve(pricePlanName); err == nil {
		fmt.Printf("[octane] Price plan \"%s\" already exists\n", pricePlanName)
		return
	}
	fmt.Printf("[octane] Price plan \"%s\" does not exist, creating\n", pricePlanName)
	createPricePlanArgs := octane.CreatePricePlanArgs{
		Name:   pricePlanName,
		Period: "month",
		MeteredComponents: []octane.MeteredComponentInputArgs{
			{
				MeterName: meterName,
				PriceScheme: &octane.PriceSchemeInputArgs{
					Prices: []octane.PriceInputArgs{
						{
							Price: float64(pricePlanRate * 100),
						},
					},
					SchemeType: "FLAT",
					UnitName:   "hour",
				},
			},
		},
	}
	if _, _, err := client.PricePlans.Create(createPricePlanArgs); err != nil {
		crash(fmt.Sprintf("Unable to create price plan: %s", err.Error()))
	}
	fmt.Printf("[octane] Price plan \"%s\" successfully created\n", pricePlanName)
}

func crash(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func envOrDefaultStr(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func envOrDefaultInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		i, _ := strconv.Atoi(val)
		return i
	}
	return fallback
}

func main() {
	checkOctaneAPIKey()
	checkOctaneResources()
	r := buildRouter()
	addr := fmt.Sprintf("%s:%d", bind, port)
	fmt.Printf("[server] Listening at http://%s/\n", addr)
	r.Run(addr)
}
