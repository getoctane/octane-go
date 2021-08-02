# Octane Go Library

[![GoDoc](https://godoc.org/github.com/getoctane/octane-go?status.svg)](https://godoc.org/github.com/getoctane/octane-go)
[![GitHub Actions status](https://github.com/getoctane/octane-go/workflows/build/badge.svg)](https://github.com/getoctane/octane-go/actions?query=workflow%3Abuild+)

[![Octane](./octane.png)](https://www.getoctane.io/)

The **[Octane](https://www.getoctane.io/)** Go library provides programmatic access
to the Octane API for Go applications.

---

- [Getting started](#getting-started)
- [Example apps](#example-apps)
- [Making API calls](#making-api-calls)
    - [Customers API](#customers-api)
        - [Example: Creating a new customer](#example-creating-a-new-customer)
        - [Example: Subscribe a customer to a price plan](#example-subscribe-a-customer-to-a-price-plan)
    - [Meters API](#meters-api)
        - [Example: Creating a new meter](#example-creating-a-new-meter)
    - [Price Plans API](#price-plans-api)
        - [Example: Creating a new price plan](#example-creating-a-new-price-plan)
    - [Measurements API](#measurements-api)
        - [Example: Sending a measurement](#example-sending-a-measurement)
- [Development](#development)
- [Contributing](#contributing)

## Getting started

First, obtain an API key from within the [Octane portal](http://cloud.getoctane.io/), and set it in your environment:

```shell
export OCTANE_API_KEY="<insert_octane_api_key_here>"
```

Then, from within your application, import the module and create a client:

```go
package main

import (
	"os"

	"github.com/getoctane/octane-go"
)

func main() {
	client := octane.NewClient(os.Getenv("OCTANE_API_KEY"))
	// ...
}
```

Update your project's `go.mod` to include the new import, and vendor the dependencies:

```shell
go mod tidy && go mod vendor
```

## Example apps

The following demo applications found in the [examples/](./examples/) directory display
how to use the Octane Go library in real-world settings:

- [antler-db-cloud-shop](examples/antler-db-cloud-shop/) - Enable your customers to self-service various cloud resources
- [customer-hours-tracker](./examples/customer-hours-tracker/) - Track hours spent working on freelance projects

## Making API calls

The `Client` type (returned by `octane.NewClient("<token>")`)
provides programmatic access to the Octane API.

### Customers API

The `Customers` namespace on a `Client` instance provides the ability to
make calls to the Octane Customers API.

#### Example: Creating a new customer

```go
customerName := "r2d2"

args := octane.CreateCustomerArgs{
    Name: customerName,
    MeasurementMappings: []octane.CustomerMeasurementMappingInputArgs{
        {
            Label:      "customer_name",
            ValueRegex: customerName,
        },
    },
}

customer, resp, err := client.Customers.Create(args)
```

#### Example: Subscribe a customer to a price plan

```go
customerName := "r2d2"

args := octane.CreateSubscriptionArgs{
    PricePlanName: "droidplan",
}

subscription, resp, err := client.Customers.CreateSubscription(customerName, args)
```

### Meters API

The `Meters` namespace on a `Client` instance provides the ability to
make calls to the Octane Meters API.

#### Example: Creating a new meter

```go
meterName := "droidrepairs"

args := octane.MeterInputArgs{
    Name:           meterName,
    MeterType:      "COUNTER",
    IsIncremental:  true,
    ExpectedLabels: []string{"customer_name"},
}

meter, resp, err := client.Meters.Create(args)
```

### Price Plans API

The `PricePlans` namespace on a `Client` instance provides the ability to
make calls to the Octane Price Plans API.

#### Example: Creating a new price plan

```go
pricePlanName := "droidplan"
pricePlanRate := 10000 // $100.00
meterName := "droidrepairs"

args := octane.CreatePricePlanArgs{
    Name:   pricePlanName,
    Period: "month",
    MeteredComponents: []octane.MeteredComponentInputArgs{
        {
            MeterName: meterName,
            PriceScheme: &octane.PriceSchemeInputArgs{
                Prices: []octane.PriceInputArgs{
                    {
                        Price: float64(pricePlanRate),
                    },
                },
                SchemeType: "FLAT",
            },
        },
    },
}

pricePlan, resp, err := client.PricePlans.Create(args)
```

### Measurements API

The `Measurements` namespace on a `Client` instance provides the ability to
make calls to the Octane Measurements API.

#### Example: Sending a measurement

```go
meterName := "droidrepairs"
customerName := "r2d2"

args := octane.Measurement{
    MeterName: meterName,
    Value: 1,
    Labels: map[string]string{
        "customer_name": customerName,
    },
}

measurement, resp, err := client.Measurements.Create(args)
```

## Development

In the root of this repo, download required dependencies:

```shell
go mod vendor
```

To regenerate files in `internal/swagger/` from latest
Octane OpenAPI spec, run the following (requires Docker):

```shell
make codegen
```

These files should then be checked into git:

```shell
git add internal/swagger/
```

## Contributing

Contributions are welcome!

Prior to submitting a 
[pull request](https://github.com/getoctane/octane-go/pulls), please
check the list of [open issues](https://github.com/getoctane/octane-go/issues).
If there is not an existing issue related to your changes, please open a
new issue to first discuss your thoughts with the project maintainers.
