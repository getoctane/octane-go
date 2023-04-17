package octane

import (
	"context"
	"net/http"

	"github.com/getoctane/octane-go/internal/swagger"
)

type (
	// Client is a client for accessing the Octane API.
	Client struct {
		Customers      *customersAPI
		Measurements   *measurementsAPI
		Meters         *metersAPI
		PricePlans     *pricePlansAPI
		CustomerPortal *customerPortalAPI
	}

	// ClientOption is used to customize a Client.
	ClientOption func(*clientSettings)

	clientSettings struct {
		basePath      string
		host          string
		scheme        string
		defaultHeader map[string]string
		userAgent     string
		httpClient    *http.Client
	}
)

// NewClient returns a new client for the Octane API.
func NewClient(token string, options ...ClientOption) *Client {
	implConfig := swagger.NewConfiguration()
	settings := &clientSettings{
		basePath:      implConfig.BasePath,
		host:          implConfig.Host,
		scheme:        implConfig.Scheme,
		defaultHeader: implConfig.DefaultHeader,
		userAgent:     implConfig.UserAgent,
		httpClient:    implConfig.HTTPClient,
	}
	for _, option := range options {
		option(settings)
	}
	impl := swagger.NewAPIClient(&swagger.Configuration{
		BasePath:      settings.basePath,
		Host:          settings.host,
		Scheme:        settings.scheme,
		DefaultHeader: settings.defaultHeader,
		UserAgent:     settings.userAgent,
		HTTPClient:    settings.httpClient,
	})
	ctx := func() context.Context {
		return context.WithValue(context.Background(),
			swagger.ContextAccessToken, token)
	}
	client := Client{
		Customers:      newCustomersAPI(impl, ctx),
		Measurements:   newMeasurementsAPI(impl, ctx),
		Meters:         newMetersAPI(impl, ctx),
		PricePlans:     newPricePlansAPI(impl, ctx),
		CustomerPortal: newCustomerPortalAPI(impl, ctx),
	}
	return &client
}

func ClientOptWithBasePath(basePath string) ClientOption {
	return func(settings *clientSettings) {
		settings.basePath = basePath
	}
}

func ClientOptWithHost(host string) ClientOption {
	return func(settings *clientSettings) {
		settings.host = host
	}
}

func ClientOptWithScheme(scheme string) ClientOption {
	return func(settings *clientSettings) {
		settings.scheme = scheme
	}
}

func ClientOptWithDefaultHeader(defaultHeader map[string]string) ClientOption {
	return func(settings *clientSettings) {
		settings.defaultHeader = defaultHeader
	}
}

func ClientOptWithUserAgent(userAgent string) ClientOption {
	return func(settings *clientSettings) {
		settings.userAgent = userAgent
	}
}

func ClientOptWithHTTPClient(httpClient *http.Client) ClientOption {
	return func(settings *clientSettings) {
		settings.httpClient = httpClient
	}
}
