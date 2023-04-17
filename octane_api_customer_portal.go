package octane

import (
	"context"
	"net/http"

	"github.com/getoctane/octane-go/internal/swagger"
)

type (
	CustomerPortalTokenInputArgs        swagger.CustomerPortalTokenInputArgs
	CustomerPortalSubscriptionInputArgs swagger.CustomerPortalSubscriptionInputArgs
	CustomerPortalToken                 swagger.CustomerPortalToken
	CustomerPortalSubscription          swagger.CustomerPortalSubscription
	CustomerPaymentMethodStatus         swagger.CustomerPaymentMethodStatus
	CustomerPortalStripeCredential      swagger.CustomerPortalStripeCredential

	customerPortalAPI struct {
		impl *swagger.APIClient
		ctx  func() context.Context
	}
)

func newCustomerPortalAPI(impl *swagger.APIClient, ctx func() context.Context) *customerPortalAPI {
	return &customerPortalAPI{impl: impl, ctx: ctx}
}

func (api *customerPortalAPI) EcpTokenPost(args CustomerPortalTokenInputArgs) (CustomerPortalToken, *http.Response, error) {
	implCustomerPortalArgs := swagger.CustomerPortalTokenInputArgs{
		CustomerName: args.CustomerName,
	}
	customerPortalTokenImpl, httpResp, err := api.impl.CustomerPortalApi.EcpTokenPost(api.ctx(), implCustomerPortalArgs)
	customerPortalToken := CustomerPortalToken{
		Token: customerPortalTokenImpl.Token,
	}
	return customerPortalToken, httpResp, err
}

func (api *customerPortalAPI) EcpSubscriptionGet(ctx context.Context) (CustomerPortalSubscription, *http.Response, error) {
	customerPortalSubscriptionImpl, httpResp, err := api.impl.CustomerPortalApi.EcpSubscriptionGet(ctx)
	customerPortalSubscription := CustomerPortalSubscription{
		PricePlan: customerPortalSubscriptionImpl.PricePlan,
	}
	return customerPortalSubscription, httpResp, err
}

func (api *customerPortalAPI) EcpPaymentMethodStatusGet(ctx context.Context) (CustomerPaymentMethodStatus, *http.Response, error) {
	customerPaymentMethodStatusImpl, httpResp, err := api.impl.CustomerPortalApi.EcpPaymentMethodStatusGet(ctx)
	customerPaymentMethodStatus := CustomerPaymentMethodStatus{
		Status: customerPaymentMethodStatusImpl.Status,
	}
	return customerPaymentMethodStatus, httpResp, err
}

func (api *customerPortalAPI) EcpPricePlansGet(ctx context.Context) ([]PricePlan, *http.Response, error) {
	implPlans, httpResp, err := api.impl.CustomerPortalApi.EcpPricePlansGet(ctx)
	plans := make([]PricePlan, 0)
	for _, plan := range implPlans {
		convertedPlan := implPricePlanToPricePlan(plan)
		plans = append(plans, convertedPlan)
	}
	return plans, httpResp, err
}

func (api *customerPortalAPI) EcpSetupIntentPost(ctx context.Context) (CustomerPortalStripeCredential, *http.Response, error) {
	customerPortalStripeCredentialImpl, httpResp, err := api.impl.CustomerPortalApi.EcpSetupIntentPost(ctx)
	customerPortalStripeCredential := CustomerPortalStripeCredential{
		ClientSecret:   customerPortalStripeCredentialImpl.ClientSecret,
		AccountId:      customerPortalStripeCredentialImpl.AccountId,
		PublishableKey: customerPortalStripeCredentialImpl.PublishableKey,
	}
	return customerPortalStripeCredential, httpResp, err
}

func (api *customerPortalAPI) EcpSubscriptionDelete(ctx context.Context) (*http.Response, error) {
	httpResp, err := api.impl.CustomerPortalApi.EcpSubscriptionDelete(ctx)
	return httpResp, err
}

func (api *customerPortalAPI) EcpSubscriptionPost(ctx context.Context, body CustomerPortalSubscriptionInputArgs) (CustomerPortalSubscription, *http.Response, error) {
	implCustomerPortalSubscriptionInputArgs := swagger.CustomerPortalSubscriptionInputArgs{
		PricePlanName: body.PricePlanName,
	}
	customerPortalSubscriptionImpl, httpResp, err := api.impl.CustomerPortalApi.EcpSubscriptionPost(ctx, implCustomerPortalSubscriptionInputArgs)
	customerPortalSubscription := CustomerPortalSubscription{
		PricePlan: customerPortalSubscriptionImpl.PricePlan,
	}
	return customerPortalSubscription, httpResp, err
}
