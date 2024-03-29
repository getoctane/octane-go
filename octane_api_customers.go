package octane

import (
	"context"
	"net/http"
	"time"

	"github.com/getoctane/octane-go/internal/swagger"
)

type (
	CreateCustomerArgs struct {
		VendorId                           int32                                 `json:"vendor_id,omitempty"`
		Tags                               []string                              `json:"tags,omitempty"`
		DisplayName                        string                                `json:"display_name,omitempty"`
		ContactInfo                        *ContactInfoInputArgs                 `json:"contact_info,omitempty"`
		Name                               string                                `json:"name,omitempty"`
		MeasurementMappings                []CustomerMeasurementMappingInputArgs `json:"measurement_mappings,omitempty"`
		PricePlanName                      string                                `json:"price_plan_name,omitempty"`
		PricePlanTag                       string                                `json:"price_plan_tag,omitempty"`
		AutogeneratePaymentGatewayCustomer bool                                  `json:"autogenerate_payment_gateway_customer,omitempty"`
		CreatedAt                          time.Time                             `json:"created_at,omitempty"`
		CustomerMetadata                   []CustomerMetadata
	}
	ContactInfoInputArgs struct {
		Country      string `json:"country,omitempty"`
		AddressLine1 string `json:"address_line_1,omitempty"`
		State        string `json:"state,omitempty"`
		Zipcode      string `json:"zipcode,omitempty"`
		Email        string `json:"email,omitempty"`
		AddressLine2 string `json:"address_line_2,omitempty"`
		City         string `json:"city,omitempty"`
		Url          string `json:"url,omitempty"`
		Phone        string `json:"phone,omitempty"`
		LegalName    string `json:"legal_name,omitempty"`
	}

	CustomerMeasurementMappingInputArgs struct {
		Label      string `json:"label,omitempty"`
		ValueRegex string `json:"value_regex,omitempty"`
	}
	CustomerMetadata struct {
		Property   string `json:"property,omitempty"`
		Value      string `json:"value,omitempty"`
	}

	UpdateCustomerArgs struct {
		VendorId            int32                                 `json:"vendor_id,omitempty"`
		Tags                []string                              `json:"tags,omitempty"`
		DisplayName         string                                `json:"display_name,omitempty"`
		ContactInfo         *ContactInfoInputArgs                 `json:"contact_info,omitempty"`
		Name                string                                `json:"name,omitempty"`
		MeasurementMappings []CustomerMeasurementMappingInputArgs `json:"measurement_mappings,omitempty"`
	}
	Customer swagger.Customer

	CustomerPaymentGatewayCredentialInputArgs swagger.CustomerPaymentGatewayCredentialInputArgs
	PaymentGatewayCredential                  swagger.PaymentGatewayCredential

	Discount struct {
		Amount float64`json:"price_plan_name,omitempty"`
		DiscountType string `json:"price_plan_name,omitempty"`
	}

	CreateSubscriptionArgs struct {
		PricePlanName      string             `json:"price_plan_name,omitempty"`
		Discounts 		   []Discount 	      `json:"discounts,omitempty"`
		CouponOverrideName string             `json:"coupon_override_name,omitempty"`
		EffectiveAt        time.Time          `json:"effective_at,omitempty"`
		CouponOverrideId   int32              `json:"coupon_override_id,omitempty"`
		CustomerId         int32              `json:"customer_id,omitempty"`
		PricePlanId        int32              `json:"price_plan_id,omitempty"`
		VendorId           int32              `json:"vendor_id,omitempty"`
		PricePlanTag       string             `json:"price_plan_tag,omitempty"`
	}

	UpdateSubscriptionArgs swagger.UpdateSubscriptionArgs
	DeleteSubscriptionArgs swagger.DeleteSubscriptionArgs
	Subscription           swagger.Subscription

	CustomerMeasurmentMapping swagger.CustomerMeasurementMapping

	CustomersApiCustomersCustomerNameRevenueGetOpts swagger.CustomersApiCustomersCustomerNameRevenueGetOpts
	RevenueResponse                                 swagger.RevenueResponse

	RevenueBreakdown struct {
		TotalRevenue float64     `json:"total_revenue,omitempty"`
		LineItems    []LineItems `json:"line_items,omitempty"`
	}
	LineItems swagger.LineItems

	CustomersApiCustomersCustomerNameUsageGetOpts swagger.CustomersApiCustomersCustomerNameUsageGetOpts
	CustomerUsage                                 swagger.CustomerUsage
	CustomerPortalTokenInputArgs                  swagger.CustomerPortalTokenInputArgs
	CustomerPortalToken                           swagger.CustomerPortalToken

	customersAPI struct {
		impl *swagger.APIClient
		ctx  func() context.Context
	}
)

func newCustomersAPI(impl *swagger.APIClient, ctx func() context.Context) *customersAPI {
	return &customersAPI{impl: impl, ctx: ctx}
}

// Create creates a new customer.
func (api *customersAPI) Create(body CreateCustomerArgs) (Customer, *http.Response, error) {
	var implMeasurementMappingInputArgs []swagger.CustomerMeasurementMappingInputArgs
	for _, mm := range body.MeasurementMappings {
		implMeasurementMappingInputArgs = append(implMeasurementMappingInputArgs, swagger.CustomerMeasurementMappingInputArgs{
			Label:      mm.Label,
			ValueRegex: mm.ValueRegex,
		})
	}
	var implContactInfo swagger.ContactInfoInputArgs
	if body.ContactInfo != nil {
		implContactInfo = swagger.ContactInfoInputArgs{
			Country:      body.ContactInfo.Country,
			AddressLine1: body.ContactInfo.AddressLine1,
			State:        body.ContactInfo.State,
			Zipcode:      body.ContactInfo.Zipcode,
			Email:        body.ContactInfo.Email,
			AddressLine2: body.ContactInfo.AddressLine2,
			City:         body.ContactInfo.City,
			Url:          body.ContactInfo.Url,
			Phone:        body.ContactInfo.Phone,
			LegalName:    body.ContactInfo.LegalName,
		}
	}
	createdAt := body.CreatedAt
	if createdAt.IsZero()  {
		createdAt = time.Now()
	}
	var newMetadata []swagger.CustomerMetadataInput
	if body.CustomerMetadata != nil {
		for _, currMetadum:= range(body.CustomerMetadata) {
			newMetadata = append(newMetadata, swagger.CustomerMetadataInput {
				Property: currMetadum.Property,
				Value: currMetadum.Value,
			})
		}
	}
	implCreateCustomerArgs := swagger.CreateCustomerArgs{
		VendorId:                           body.VendorId,
		Tags:                               body.Tags,
		DisplayName:                        body.DisplayName,
		ContactInfo:                        &implContactInfo,
		Name:                               body.Name,
		MeasurementMappings:                implMeasurementMappingInputArgs,
		PricePlanName:                      body.PricePlanName,
		PricePlanTag:                       body.PricePlanTag,
		AutogeneratePaymentGatewayCustomer: body.AutogeneratePaymentGatewayCustomer,
		CreatedAt:                          createdAt,
		CustomerMetadata:                   newMetadata,
	}
	implCustomer, resp, err := api.impl.CustomersApi.CustomersPost(
		api.ctx(), implCreateCustomerArgs)
	customer := implCustomerToCustomer(&implCustomer)
	return customer, resp, err
}

// Retrieve fetches a customer by their unique name.
func (api *customersAPI) Retrieve(customerName string) (Customer, *http.Response, error) {
	implCustomer, resp, err := api.impl.CustomersApi.CustomersCustomerNameGet(
		api.ctx(), customerName)
	customer := implCustomerToCustomer(&implCustomer)
	return customer, resp, err
}

// Update updates a single customer.
func (api *customersAPI) Update(customerName string, body UpdateCustomerArgs) (Customer, *http.Response, error) {
	var implMeasurementMappingInputArgs []swagger.CustomerMeasurementMappingInputArgs
	for _, mm := range body.MeasurementMappings {
		implMeasurementMappingInputArgs = append(implMeasurementMappingInputArgs, swagger.CustomerMeasurementMappingInputArgs{
			Label:      mm.Label,
			ValueRegex: mm.ValueRegex,
		})
	}
	var implContactInfo swagger.ContactInfoInputArgs
	if body.ContactInfo != nil {
		implContactInfo = swagger.ContactInfoInputArgs{
			Country:      body.ContactInfo.Country,
			AddressLine1: body.ContactInfo.AddressLine1,
			State:        body.ContactInfo.State,
			Zipcode:      body.ContactInfo.Zipcode,
			Email:        body.ContactInfo.Email,
			AddressLine2: body.ContactInfo.AddressLine2,
			City:         body.ContactInfo.City,
			Url:          body.ContactInfo.Url,
			Phone:        body.ContactInfo.Phone,
			LegalName:    body.ContactInfo.LegalName,
		}
	}
	implUpdateCustomerArgs := swagger.UpdateCustomerArgs{
		VendorId:            body.VendorId,
		Tags:                body.Tags,
		DisplayName:         body.DisplayName,
		ContactInfo:         &implContactInfo,
		Name:                body.Name,
		MeasurementMappings: implMeasurementMappingInputArgs,
	}
	implCustomer, resp, err := api.impl.CustomersApi.CustomersCustomerNamePut(
		api.ctx(), implUpdateCustomerArgs, customerName)
	customer := implCustomerToCustomer(&implCustomer)
	return customer, resp, err
}

// List retrieves all customers for a given vendor.
func (api *customersAPI) List() ([]Customer, *http.Response, error) {
	implCustomers, resp, err := api.impl.CustomersApi.CustomersGet(api.ctx())
	var customers []Customer
	for _, implCustomer := range implCustomers {
		customers = append(customers, implCustomerToCustomer(&implCustomer))
	}
	return customers, resp, err
}

// Delete deletes a customer by their unique name.
func (api *customersAPI) Delete(customerName string) (*http.Response, error) {
	return api.impl.CustomersApi.CustomersCustomerNameDelete(
		api.ctx(), customerName)
}

// CreatePaymentGatewayCredential adds PaymentGatewayCredential for a customer.
func (api *customersAPI) CreatePaymentGatewayCredential(customerName string, body CustomerPaymentGatewayCredentialInputArgs) (PaymentGatewayCredential, *http.Response, error) {
	implCustomerPaymentGatewayCredentialInputArgs := swagger.CustomerPaymentGatewayCredentialInputArgs{
		AccountId: body.AccountId,
	}
	implPaymentGatewayCredential, resp, err := api.impl.CustomersApi.CustomersCustomerNamePaymentGatewayCredentialsPost(
		api.ctx(), implCustomerPaymentGatewayCredentialInputArgs, customerName)
	paymentGatewayCredential := implPaymentGatewayCredentialToPaymentGatewayCredential(&implPaymentGatewayCredential)
	return paymentGatewayCredential, resp, err
}

// CreateSubscription creates a new subscription for a customer / price plan combination (by unique name).
func (api *customersAPI) CreateSubscription(customerName string, body CreateSubscriptionArgs) (Subscription, *http.Response, error) {
	implCreateSubscriptionArgs := swagger.CreateSubscriptionArgs{
		VendorId:           body.VendorId,
		CouponOverrideId:   body.CouponOverrideId,
		CustomerId:         body.CustomerId,
		EffectiveAt:        body.EffectiveAt,
		PricePlanName:      body.PricePlanName,
		CouponOverrideName: body.CouponOverrideName,
		PricePlanId:        body.PricePlanId,
		PricePlanTag:       body.PricePlanTag,
	}
	if body.Discounts != nil {
		var discounts []swagger.DiscountInputArgs
		for _, implDiscount := range body.Discounts {
			discounts = append(discounts, swagger.DiscountInputArgs{
												Amount:       implDiscount.Amount,
												DiscountType: implDiscount.DiscountType,
											})
		}
		implCreateSubscriptionArgs.Discounts = discounts
	}
	if implCreateSubscriptionArgs.EffectiveAt.IsZero() {
		implCreateSubscriptionArgs.EffectiveAt = time.Now()
	}
	implSubscription, resp, err := api.impl.CustomersApi.CustomersCustomerNameSubscriptionsPost(
		api.ctx(), implCreateSubscriptionArgs, customerName)
	subscription := implSubscriptionToSubscription(&implSubscription)
	return subscription, resp, err
}

// ListSubscriptions gets all subscriptions for the customer.
func (api *customersAPI) ListSubscriptions(customerName string) ([]Subscription, *http.Response, error) {
	implSubscriptions, resp, err := api.impl.CustomersApi.CustomersCustomerNameSubscriptionsGet(
		api.ctx(), customerName)
	var subscriptions []Subscription
	for _, implSubscription := range implSubscriptions {
		subscriptions = append(subscriptions, implSubscriptionToSubscription(&implSubscription))
	}
	return subscriptions, resp, err
}

// UpdateSubscription updates a subscription for a specific customer (by customer name).
func (api *customersAPI) UpdateSubscription(customerName string, body UpdateSubscriptionArgs) (Subscription, *http.Response, error) {
	implUpdateSubscriptionArgs := swagger.UpdateSubscriptionArgs{
		VendorId:           body.VendorId,
		CouponOverrideId:   body.CouponOverrideId,
		CustomerId:         body.CustomerId,
		DiscountOverride:   body.DiscountOverride,
		EffectiveAt:        body.EffectiveAt,
		PricePlanName:      body.PricePlanName,
		CouponOverrideName: body.CouponOverrideName,
		PricePlanId:        body.PricePlanId,
	}
	implSubscription, resp, err := api.impl.CustomersApi.CustomersCustomerNameSubscriptionPut(
		api.ctx(), implUpdateSubscriptionArgs, customerName)
	subscription := implSubscriptionToSubscription(&implSubscription)
	return subscription, resp, err
}

// DeleteSubscription deletes a subscription for a specific customer (by customer name).
func (api *customersAPI) DeleteSubscription(customerName string, body DeleteSubscriptionArgs) (*http.Response, error) {
	implDeleteSubscriptionArgs := swagger.DeleteSubscriptionArgs{
		ExpireAt:   body.ExpireAt,
		VendorId:   body.VendorId,
		CustomerId: body.CustomerId,
	}
	return api.impl.CustomersApi.CustomersCustomerNameSubscriptionDelete(
		api.ctx(), implDeleteSubscriptionArgs, customerName)
}

// CreateMapping creates a measurement mapping for a speciifc customer (by customer name).
func (api *customersAPI) CreateMapping(customerName string, body CustomerMeasurementMappingInputArgs) (CustomerMeasurmentMapping, *http.Response, error) {
	implCustomerMeasurementMappingInputArgs := swagger.CustomerMeasurementMappingInputArgs{
		ValueRegex: body.ValueRegex,
		Label:      body.Label,
	}
	implCustomerMeasurmentMapping, resp, err := api.impl.CustomersApi.CustomersCustomerNameMappingsPost(
		api.ctx(), implCustomerMeasurementMappingInputArgs, customerName)
	customerMeasurmentMapping := implCustomerMeasurmentMappingToCustomerMeasurmentMapping(&implCustomerMeasurmentMapping)
	return customerMeasurmentMapping, resp, err
}

// RetrieveRevenue fetches a customer revenue data by their unique name.
func (api *customersAPI) RetrieveRevenue(customerName string, body CustomersApiCustomersCustomerNameRevenueGetOpts) (RevenueResponse, *http.Response, error) {
	implCustomersApiCustomersCustomerNameRevenueGetOpts := swagger.CustomersApiCustomersCustomerNameRevenueGetOpts{
		StartTime: body.StartTime,
		EndTime:   body.EndTime,
	}
	implRevenueResponse, resp, err := api.impl.CustomersApi.CustomersCustomerNameRevenueGet(
		api.ctx(), customerName, &implCustomersApiCustomersCustomerNameRevenueGetOpts)
	revenueResponse := implRevenueResponseToRevenueResponse(&implRevenueResponse)
	return revenueResponse, resp, err
}

// RetrieveAccruedRevenue fetches a customer accrued revenue data by their unique name.
func (api *customersAPI) RetrieveAccruedRevenue(customerName string) (RevenueBreakdown, *http.Response, error) {
	implRevenueBreakdown, resp, err := api.impl.CustomersApi.CustomersCustomerNameAccruedRevenueGet(
		api.ctx(), customerName)
	revenueBreakdown := implRevenueBreakdownToRevenueBreakdown(&implRevenueBreakdown)
	return revenueBreakdown, resp, err
}

// RetrieveUsage fetches a customer usage data by their unique name.
func (api *customersAPI) RetrieveUsage(customerName string, body CustomersApiCustomersCustomerNameUsageGetOpts) (CustomerUsage, *http.Response, error) {
	implCustomersApiCustomersCustomerNameUsageGetOpts := swagger.CustomersApiCustomersCustomerNameUsageGetOpts{
		MeterName: body.MeterName,
		StartTime: body.StartTime,
		EndTime:   body.EndTime,
	}
	implCustomerUsage, resp, err := api.impl.CustomersApi.CustomersCustomerNameUsageGet(
		api.ctx(), customerName, &implCustomersApiCustomersCustomerNameUsageGetOpts)
	customerUsage := implCustomerUsageToCustomerUsage(&implCustomerUsage)
	return customerUsage, resp, err
}

func (api *customersAPI) GenerateEcpToken(args CustomerPortalTokenInputArgs) (CustomerPortalToken, *http.Response, error) {
	implCustomerPortalArgs := swagger.CustomerPortalTokenInputArgs{
		CustomerName: args.CustomerName,
	}
	customerPortalTokenImpl, httpResp, err := api.impl.CustomerPortalApi.EcpTokenPost(api.ctx(), implCustomerPortalArgs)
	customerPortalToken := CustomerPortalToken{
		Token: customerPortalTokenImpl.Token,
	}
	return customerPortalToken, httpResp, err
}

// Convert a Swagger Customer struct to our Customer struct
func implCustomerToCustomer(implCustomer *swagger.Customer) Customer {
	var customer Customer
	if implCustomer != nil {
		customer = Customer{
			Name:                implCustomer.Name,
			DisplayName:         implCustomer.DisplayName,
			ContactInfo:         implCustomer.ContactInfo,
			MeasurementMappings: implCustomer.MeasurementMappings,
		}
	}
	return customer
}

// Convert a Swagger Subscription struct to our Subscription struct
func implSubscriptionToSubscription(implSubscription *swagger.Subscription) Subscription {
	var subscription Subscription
	if implSubscription != nil {
		subscription = Subscription{
			CustomerName:      implSubscription.CustomerName,
			PricePlanName:     implSubscription.PricePlanName,
			TrialOverride:     implSubscription.TrialOverride,
			BasePriceOverride: implSubscription.BasePriceOverride,
			EffectiveAt:       implSubscription.EffectiveAt,
			Discounts:         implSubscription.Discounts,
		}
	}
	return subscription
}

// Convert a Swagger PaymentGatewayCredential struct to our PaymentGatewayCredential struct
func implPaymentGatewayCredentialToPaymentGatewayCredential(implPaymentGatewayCredential *swagger.PaymentGatewayCredential) PaymentGatewayCredential {
	var paymentGatewayCredential PaymentGatewayCredential
	if implPaymentGatewayCredential != nil {
		paymentGatewayCredential = PaymentGatewayCredential{
			PaymentGateway: implPaymentGatewayCredential.PaymentGateway,
			AccountId:      implPaymentGatewayCredential.AccountId,
			AuthToken:      implPaymentGatewayCredential.AuthToken,
		}
	}
	return paymentGatewayCredential
}

// Convert a Swagger CustomerMeasurmentMapping struct to our CustomerMeasurmentMapping struct
func implCustomerMeasurmentMappingToCustomerMeasurmentMapping(implCustomerMeasurmentMapping *swagger.CustomerMeasurementMapping) CustomerMeasurmentMapping {
	return CustomerMeasurmentMapping{
		Label:      implCustomerMeasurmentMapping.Label,
		ValueRegex: implCustomerMeasurmentMapping.ValueRegex,
	}
}

// Convert a Swagger RevenueResponse struct to our RevenueResponse struct
func implRevenueResponseToRevenueResponse(implRevenueResponse *swagger.RevenueResponse) RevenueResponse {
	return RevenueResponse{
		Revenue: implRevenueResponse.Revenue,
	}
}

// Convert a Swagger RevenueBreakdown struct to our RevenueBreakdown struct
func implRevenueBreakdownToRevenueBreakdown(implRevenueBreakdown *swagger.RevenueBreakdown) RevenueBreakdown {
	var lineItems []LineItems
	for _, implLineItem := range implRevenueBreakdown.LineItems {
		lineItems = append(lineItems, LineItems{
			PriceInt:     implLineItem.PriceInt,
			Description:  implLineItem.Description,
			QuantityUnit: implLineItem.QuantityUnit,
			Name:         implLineItem.Name,
			Quantity:     implLineItem.Quantity,
			Price:        implLineItem.Price,
			Id:           implLineItem.Id,
			Metadata:     implLineItem.Metadata,
		})
	}
	return RevenueBreakdown{
		TotalRevenue: implRevenueBreakdown.TotalRevenue,
		LineItems:    lineItems,
	}
}

// Convert a Swagger CustomerUsage struct to our CustomerUsage struct
func implCustomerUsageToCustomerUsage(implCustomerUsage *swagger.CustomerUsage) CustomerUsage {
	return CustomerUsage{
		Usage: implCustomerUsage.Usage,
	}
}
