#!/usr/bin/env bash

set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR/../

OPEN_API_URL="https://api.cloud.getoctane.io/docs/openapi.json"

CODEGEN_VERSION="3.0.27"
CODEGEN_SHA="66e956839d84bfff44be2ac269761f755404dfea34c7e4903821fedbc3c06043"
CODEGEN_URL="https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/${CODEGEN_VERSION}/swagger-codegen-cli-${CODEGEN_VERSION}.jar"

mkdir -p testbin/

if [ ! -f "testbin/openapi.json" ]; then
  curl -L "${OPEN_API_URL}" -o "testbin/openapi.json"
fi

if [ ! -f "testbin/codegen.jar" ]; then
  mkdir -p testbin/
  curl -L "${CODEGEN_URL}" -o "testbin/codegen.jar"
fi

shasum -a 256 "testbin/codegen.jar" | grep "^${CODEGEN_SHA}  "

rm -rf mount/
trap "rm -rf ${PWD}/mount" EXIT
mkdir -p mount/

# Generate the typescript files from openapi spec
docker run --rm \
  -v ${PWD}/mount:/mount \
  -v ${PWD}/testbin:/testbin \
  --entrypoint java \
  java:8 \
  -jar /testbin/codegen.jar generate \
  -i /testbin/openapi.json -l go -o /mount

# Fixes for strange "Object" types
cat mount/model_all_of_subscription_discount_override.go | \
  sed 's/DiscountType \*Object/DiscountType string/g' \
  > mount/model_all_of_subscription_discount_override.go.tmp
mv mount/model_all_of_subscription_discount_override.go.tmp \
  mount/model_all_of_subscription_discount_override.go

cat mount/model_billing_settings.go | \
  sed 's/CustomerInvoiceDetailLevel \*Object/CustomerInvoiceDetailLevel string/g' \
  > mount/model_billing_settings.go.tmp
mv mount/model_billing_settings.go.tmp \
  mount/model_billing_settings.go

cat mount/model_customer.go | \
  sed 's/MeasurementMappings \[\]Object/MeasurementMappings \[\]map\[string\]string/g' \
  > mount/model_customer.go.tmp
mv mount/model_customer.go.tmp \
  mount/model_customer.go

cat mount/model_discount.go | \
  sed 's/DiscountType \*Object/DiscountType string/g' \
  > mount/model_discount.go.tmp
mv mount/model_discount.go.tmp \
  mount/model_discount.go

cat mount/model_meter.go | \
  sed 's/MeterType \*Object/MeterType string/g' | \
  sed 's/UnitName \*Object/UnitName string/g' | \
  sed 's/ExpectedLabels \[\]Object/ExpectedLabels \[\]string/g' | \
  sed 's/PrimaryLabels \[\]Object/PrimaryLabels \[\]string/g' \
  > mount/model_meter.go.tmp
mv mount/model_meter.go.tmp \
  mount/model_meter.go

cat mount/model_metered_component.go | \
  sed 's/MeterName \*Object/MeterName string/g' \
  > mount/model_metered_component.go.tmp
mv mount/model_metered_component.go.tmp \
  mount/model_metered_component.go

cat mount/model_price_scheme.go | \
  sed 's/SchemeType \*Object/SchemeType string/g' \
  > mount/model_price_scheme.go.tmp
mv mount/model_price_scheme.go.tmp \
  mount/model_price_scheme.go

cat mount/model_subscription.go | \
  sed 's/CustomerName \*Object/CustomerName string/g' | \
  sed 's/PricePlanName \*Object/PricePlanName string/g' \
  > mount/model_subscription.go.tmp
mv mount/model_subscription.go.tmp \
  mount/model_subscription.go

cat mount/model_active_subscription.go | \
  sed 's/CustomerName \*Object/CustomerName string/g' | \
  sed 's/PricePlanName \*Object/PricePlanName string/g' \
  > mount/model_active_subscription.go.tmp
mv mount/model_active_subscription.go.tmp \
  mount/model_active_subscription.go

cat mount/model_all_of_active_subscription_discount_override.go | \
  sed 's/DiscountType \*Object/DiscountType string/g' \
  > mount/model_all_of_active_subscription_discount_override.go.tmp
mv mount/model_all_of_active_subscription_discount_override.go.tmp \
  mount/model_all_of_active_subscription_discount_override.go

cat mount/model_payment_gateway_credential.go | \
  sed 's/PaymentGateway \*Object/PaymentGateway string/g' \
  > mount/model_payment_gateway_credential.go.tmp
mv mount/model_payment_gateway_credential.go.tmp \
  mount/model_payment_gateway_credential.go

rm -f internal/swagger/*.go
mkdir -p internal/swagger/
cp mount/*.go internal/swagger/
