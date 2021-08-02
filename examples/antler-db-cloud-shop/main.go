package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/getoctane/octane-go"
	"github.com/gin-gonic/gin"
)

var (
	// Instantiate the Octane API client
	client = octane.NewClient(os.Getenv("OCTANE_API_KEY"))

	// Application settings and defaults
	port               = envOrDefaultInt("APP_PORT", 3000)
	bind               = envOrDefaultStr("APP_BIND", "127.0.0.1")
	octaneRedirectUrl  = envOrDefaultStr("OCTANE_REDIRECT_URL", "https://cloud.getoctane.io")
	pricePlanName      = envOrDefaultStr("OCTANE_PRICE_PLAN_NAME", "antlerdb")
	meterNameStorage   = envOrDefaultStr("OCTANE_METER_NAME_STORAGE", "storage")
	meterNameBandwidth = envOrDefaultStr("OCTANE_METER_NAME_BANDWIDTH", "bandwidth")
	meterNameMachines  = envOrDefaultStr("OCTANE_METER_NAME_MACHINES", "machines")
	meterRateStorage   = envOrDefaultInt("OCTANE_METER_RATE_STORAGE", 2)
	meterRateBandwidth = envOrDefaultInt("OCTANE_METER_RATE_BANDWIDTH", 5)
	meterRateMachines  = envOrDefaultInt("OCTANE_METER_RATE_MACHINES", 10)

	// The frontend sends us generic resource names,
	// which we convert to meter names
	resourceMeterMap = map[string]string{
		"storage":   meterNameStorage,
		"bandwidth": meterNameBandwidth,
		"machines":  meterNameMachines,
	}
)

func buildRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFile("/", "./public/index.html")
	r.StaticFile("/index.html", "./public/index.html")
	r.StaticFile("/scripts.js", "./public/scripts.js")
	r.StaticFile("/styles.css", "./public/styles.css")
	r.GET("/api/whoami", handleGetAPIWhoami)
	r.POST("/api/resources", handlerPostAPIResources)
	return r
}

func handleGetAPIWhoami(c *gin.Context) {
	if name, err := c.Cookie("username"); err == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"name": name,
			"url":  octaneRedirectUrl,
		})
		return
	}
	name := petname.Generate(2, "-")
	fmt.Printf("[octane] Attempting to create new customer \"%s\"\n", name)
	createCustomerArgs := octane.CreateCustomerArgs{
		Name: name,
		MeasurementMappings: []octane.CustomerMeasurementMappingInputArgs{
			{
				Label:      "customer_name",
				ValueRegex: name,
			},
		},
	}
	if _, _, err := client.Customers.Create(createCustomerArgs); err != nil {
		fmt.Printf("[octane] Error creating customer \"%s\"\n", name)
		badRequest(c, err)
		return
	}
	fmt.Printf("[octane] Customer \"%s\" successfully created\n", name)
	fmt.Printf("[[octane] Attempting to subscribe customer \"%s\" to price plan \"%s\"\n",
		name, pricePlanName)
	createSubscriptionArgs := octane.CreateSubscriptionArgs{
		PricePlanName: pricePlanName,
	}
	if _, _, err := client.Customers.CreateSubscription(name, createSubscriptionArgs); err != nil {
		fmt.Printf(
			"[octane] Error subscribing customer \"%s\" to price plan \"%s\"\n",
			name, pricePlanName)
		badRequest(c, err)
		return
	}
	fmt.Printf(
		"[octane] Successfully subscribed customer \"%s\" to price plan \"%s\"\n",
		name, pricePlanName)
	c.SetCookie("username", name, 86400, "",
		fmt.Sprintf("%s:%d", bind, port), false, true)
	c.JSON(200, gin.H{
		"code": 200,
		"name": name,
		"url":  octaneRedirectUrl,
	})
}

func handlerPostAPIResources(c *gin.Context) {
	name, err := c.Cookie("username")
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "No session, please refresh",
		})
		return
	}
	var data struct {
		Resource string
		Value    interface{}
	}
	if err := c.BindJSON(&data); err != nil {
		badRequest(c, err)
		return
	}
	meterName, ok := resourceMeterMap[data.Resource]
	if !ok {
		badRequest(c, errors.New("invalid resource provided"))
		return
	}
	var value float64
	if v, err := strconv.ParseFloat(fmt.Sprintf("%v", data.Value), 8); err != nil {
		badRequest(c, err)
		return
	} else {
		value = v
	}
	fmt.Printf(
		"[octane] Attempting to create measurement for customer \"%s\" for meter \"%s\"\n",
		name, meterName)
	measurement := octane.Measurement{
		Value: value,
		Labels: map[string]string{
			"customer_name": name,
		},
		MeterName: meterName,
	}
	fmt.Printf("[octane] Attempting to create measurement for customer \"%s\"\n", name)
	if _, _, err := client.Measurements.Create(measurement); err != nil {
		fmt.Printf("[octane] Error creating measurement for customer \"%s\"\n", name)
		badRequest(c, err)
		return
	}
	fmt.Printf("[octane] Measurement for customer \"%s\" successfully created\n", name)
	c.JSON(201, gin.H{
		"code":    201,
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

func checkOctaneResourceMeter(meterInputArgs octane.MeterInputArgs) {
	fmt.Printf("[octane] Checking if meter \"%s\" exists\n", meterInputArgs.Name)
	if _, _, err := client.Meters.Retrieve(meterInputArgs.Name); err == nil {
		fmt.Printf("[octane] Meter \"%s\" already exists\n", meterInputArgs.Name)
		return
	}
	fmt.Printf("[octane] Meter \"%s\" does not exist, creating\n", meterInputArgs.Name)
	if _, _, err := client.Meters.Create(meterInputArgs); err != nil {
		crash(fmt.Sprintf("Unable to create meter: %s", err.Error()))
	}
	fmt.Printf("[octane] Meter \"%s\" successfully created\n", meterInputArgs.Name)
}

func checkOctaneResourceMeterStorage() {
	checkOctaneResourceMeter(octane.MeterInputArgs{
		Name:           meterNameStorage,
		DisplayName:    "Storage in gigabytes",
		MeterType:      "COUNTER",
		UnitName:       "gigabyte",
		IsIncremental:  true,
		ExpectedLabels: []string{"customer_name"},
	})
}

func checkOctaneResourceMeterBandwidth() {
	checkOctaneResourceMeter(octane.MeterInputArgs{
		Name:           meterNameBandwidth,
		DisplayName:    "Bandwidth in gigabytes",
		MeterType:      "COUNTER",
		UnitName:       "gigabyte",
		IsIncremental:  true,
		ExpectedLabels: []string{"customer_name"},
	})
}

func checkOctaneResourceMeterMachines() {
	checkOctaneResourceMeter(octane.MeterInputArgs{
		Name:           meterNameMachines,
		DisplayName:    "Number of machines",
		MeterType:      "COUNTER",
		IsIncremental:  true,
		ExpectedLabels: []string{"customer_name"},
	})
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
				MeterName: meterNameStorage,
				PriceScheme: &octane.PriceSchemeInputArgs{
					Prices: []octane.PriceInputArgs{
						{
							Price: float64(meterRateStorage * 100),
						},
					},
					SchemeType: "FLAT",
					UnitName:   "gigabyte",
				},
			},
			{
				MeterName: meterNameBandwidth,
				PriceScheme: &octane.PriceSchemeInputArgs{
					Prices: []octane.PriceInputArgs{
						{
							Price: float64(meterRateBandwidth * 100),
						},
					},
					SchemeType: "FLAT",
					UnitName:   "gigabyte",
				},
			},
			{
				MeterName: meterNameMachines,
				PriceScheme: &octane.PriceSchemeInputArgs{
					Prices: []octane.PriceInputArgs{
						{
							Price: float64(meterRateMachines * 100),
						},
					},
					SchemeType: "FLAT",
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
	checkOctaneResourceMeterStorage()
	checkOctaneResourceMeterBandwidth()
	checkOctaneResourceMeterMachines()
	checkOctaneResourcePricePlan()
	r := buildRouter()
	addr := fmt.Sprintf("%s:%d", bind, port)
	fmt.Printf("[server] Listening at http://%s/\n", addr)
	r.Run(addr)
}
