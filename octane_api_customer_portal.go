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
