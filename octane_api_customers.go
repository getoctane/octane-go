package octane

import (
	"context"
	"net/http"
	"time"

	"github.com/getoctane/octane-go/internal/swagger"
)

type (
	CreateCustomerArgs struct {
		VendorId            int32                                 `json:"vendor_id,omitempty"`
		Tags                []string                              `json:"tags,omitempty"`
		DisplayName         string                                `json:"display_name,omitempty"`
		ContactInfo         *ContactInfoInputArgs                 `json:"contact_info,omitempty"`
		Name                string                                `json:"name,omitempty"`
		MeasurementMappings []CustomerMeasurementMappingInputArgs `json:"measurement_mappings,omitempty"`
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
	}

	CustomerMeasurementMappingInputArgs struct {
		Label      string `json:"label,omitempty"`
		ValueRegex string `json:"value_regex,omitempty"`
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

	CreateSubscriptionArgs swagger.CreateSubscriptionArgs
	UpdateSubscriptionArgs swagger.UpdateSubscriptionArgs
	DeleteSubscriptionArgs swagger.DeleteSubscriptionArgs
	Subscription           swagger.Subscription

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
		}
	}
	implCreateCustomerArgs := swagger.CreateCustomerArgs{
		VendorId:            body.VendorId,
		Tags:                body.Tags,
		DisplayName:         body.DisplayName,
		ContactInfo:         &implContactInfo,
		Name:                body.Name,
		MeasurementMappings: implMeasurementMappingInputArgs,
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
		DiscountOverride:   body.DiscountOverride,
		EffectiveAt:        body.EffectiveAt,
		PricePlanName:      body.PricePlanName,
		CouponOverrideName: body.CouponOverrideName,
		PricePlanId:        body.PricePlanId,
	}
	// TODO: if EffectiveAt is nil, causes 500 error
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
			DiscountOverride:  implSubscription.DiscountOverride,
			TrialOverride:     implSubscription.TrialOverride,
			BasePriceOverride: implSubscription.BasePriceOverride,
			EffectiveAt:       implSubscription.EffectiveAt,
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
