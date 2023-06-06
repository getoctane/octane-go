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

# Generate the golang files from openapi spec
docker run --rm \
  -v ${PWD}/mount:/mount \
  -v ${PWD}/testbin:/testbin \
  --entrypoint java \
  arm64v8/openjdk \
  --add-opens=java.base/java.util=ALL-UNNAMED \
  -jar /testbin/codegen.jar generate \
  -i /testbin/openapi.json -l go -o /mount \
  

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

cat mount/model_metered_component_label_limit.go | \
  sed 's/Labels \*Object/Labels map[string]string/g' \
  > mount/model_metered_component_label_limit.go.tmp
mv mount/model_metered_component_label_limit.go.tmp \
  mount/model_metered_component_label_limit.go

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

cat mount/model_payment_gateway_credential.go | \
  sed 's/PaymentGateway \*Object/PaymentGateway string/g' \
  > mount/model_payment_gateway_credential.go.tmp
mv mount/model_payment_gateway_credential.go.tmp \
  mount/model_payment_gateway_credential.go

cat mount/model_all_of_subscription_price_plan.go | \
  sed 's/Coupon \*Object/Coupon \*Coupon/g' | \
  sed 's/Discount \*Object/Discount \*Discount/g' \
  > mount/model_all_of_subscription_price_plan.go.tmp
mv mount/model_all_of_subscription_price_plan.go.tmp \
  mount/model_all_of_subscription_price_plan.go


cat mount/model_all_of_active_subscription_price_plan.go | \
  sed 's/Coupon \*Object/Coupon \*Coupon/g' | \
  sed 's/Discount \*Object/Discount \*Discount/g' \
  > mount/model_all_of_active_subscription_price_plan.go.tmp
mv mount/model_all_of_active_subscription_price_plan.go.tmp \
  mount/model_all_of_active_subscription_price_plan.go

cat mount/model_all_of_price_plan_discount.go | \
  sed 's/DiscountType \*Object/DiscountType string/g' \
  > mount/model_all_of_price_plan_discount.go.tmp
mv mount/model_all_of_price_plan_discount.go.tmp \
  mount/model_all_of_price_plan_discount.go


cat mount/model_all_of_active_subscription_price_plan.go | \
  sed 's/Coupon \*Object/Coupon \*Coupon/g' | \
  sed 's/Discount \*Object/Discount \*Discount/g' \
  > mount/model_all_of_active_subscription_price_plan.go.tmp
mv mount/model_all_of_active_subscription_price_plan.go.tmp \
  mount/model_all_of_active_subscription_price_plan.go

cat mount/model_coupon1.go | \
  sed 's/DiscountType \*Object/DiscountType string/g' | \
  sed 's/Frequency \*Object/Frequency string/g' \
  > mount/model_coupon1.go.tmp
mv mount/model_coupon1.go.tmp \
  mount/model_coupon1.go

cat mount/model_create_price_plan_args.go | \
  sed 's/Discount \*AllOfCreatePricePlanArgsDiscount/Discount \*DiscountInputArgs/g' \
  > mount/model_create_price_plan_args.go.tmp
mv mount/model_create_price_plan_args.go.tmp \
  mount/model_create_price_plan_args.go

cat mount/model_update_price_plan_args.go | \
  sed 's/Discount \*AllOfUpdatePricePlanArgsDiscount/Discount \*DiscountInputArgs/g' \
  > mount/model_update_price_plan_args.go.tmp
mv mount/model_update_price_plan_args.go.tmp \
  mount/model_update_price_plan_args.go


#we want to force all timestamps sent to our api to conform to a certain format, so we help the caller here
#by changing what is sent over the wire
rfc_timestamp_insertion="func parameterToString(obj interface{}, collectionFormat string) string {\n\n \
\tif t, isTime := obj.(time.Time); isTime { \n\
\t\treturn t.Format(time.RFC3339) \n \
\t}\n"

cat mount/client.go | \
  sed "s/func parameterToString.*/$rfc_timestamp_insertion/" \
  > mount/client.go.tmp
mv mount/client.go.tmp \
  mount/client.go

cat mount/model_subscription.go | \
  sed 's/AddOns \*Object/AddOns \AddOn/g' \
  > mount/model_subscription.go.tmp
mv mount/model_subscription.go.tmp \
  mount/model_subscription.go


cat mount/model_credit_grant.go | \
  sed 's/ExpiresAt time.Time/ExpiresAt \string/g' \
  > mount/model_credit_grant.go.tmp
mv mount/model_credit_grant.go.tmp \
  mount/model_credit_grant.go

cat mount/model_credit_grant.go | \
  sed 's/EffectiveAt time.Time/EffectiveAt \string/g' \
  > mount/model_credit_grant.go.tmp
mv mount/model_credit_grant.go.tmp \
  mount/model_credit_grant.go


cat mount/model_all_of_active_subscription_discounts_items.go | \
  sed 's/DiscountType \*Object/DiscountType \string/g' \
  > mount/model_all_of_active_subscription_discounts_items.go.tmp
mv mount/model_all_of_active_subscription_discounts_items.go.tmp \
  mount/model_all_of_active_subscription_discounts_items.go

cat mount/model_all_of_active_subscription_discounts_items.go | \
  sed 's/MeteredComponent \*Object/MeteredComponent \MeteredComponent/g' \
  > mount/model_all_of_active_subscription_discounts_items.go.tmp
mv mount/model_all_of_active_subscription_discounts_items.go.tmp \
  mount/model_all_of_active_subscription_discounts_items.go

cat mount/model_all_of_active_subscription_discounts_items.go | \
  sed 's/AddOn \*Object/AddOn \AddOn/g' \
  > mount/model_all_of_active_subscription_discounts_items.go.tmp
mv mount/model_all_of_active_subscription_discounts_items.go.tmp \
  mount/model_all_of_active_subscription_discounts_items.go


cat mount/model_active_subscription.go | \
  sed 's/AddOns \*Object/AddOns \[]SubscriptionAddOnItem/g' \
  > mount/model_active_subscription.go.tmp
mv mount/model_active_subscription.go.tmp \
  mount/model_active_subscription.go

cat mount/model_all_of_customer_portal_usage_current_cycle_usage.go | \
  sed 's/UsageByTime \[]Object/UsageByTime\ []CycleUsage/g' \
  > mount/model_all_of_customer_portal_usage_current_cycle_usage.go.tmp
mv mount/model_all_of_customer_portal_usage_current_cycle_usage.go.tmp \
  mount/model_all_of_customer_portal_usage_current_cycle_usage.go


cat mount/model_all_of_customer_portal_usage_previous_cycle_usage.go | \
  sed 's/UsageByTime \[]Object/UsageByTime\ []CycleUsage/g' \
  > mount/model_all_of_customer_portal_usage_previous_cycle_usage.go.tmp
mv mount/model_all_of_customer_portal_usage_previous_cycle_usage.go.tmp \
  mount/model_all_of_customer_portal_usage_previous_cycle_usage.go

cat mount/model_coupon.go | \
  sed 's/Frequency \*Object/Frequency \string/g' \
  > mount/model_coupon.go.tmp
mv mount/model_coupon.go.tmp \
  mount/model_coupon.go

cat mount/model_coupon.go | \
  sed 's/DiscountType \*Object/DiscountType \string/g' \
  > mount/model_coupon.go.tmp
mv mount/model_coupon.go.tmp \
  mount/model_coupon.go

cat mount/model_all_of_subscription_discounts_items.go | \
  sed 's/DiscountType \*Object/DiscountType \string/g' \
  > mount/model_all_of_subscription_discounts_items.go.tmp
mv mount/model_all_of_subscription_discounts_items.go.tmp \
  mount/model_all_of_subscription_discounts_items.go

cat mount/model_all_of_subscription_discounts_items.go | \
  sed 's/AddOn \*Object/AddOn \AddOn/g' \
  > mount/model_all_of_subscription_discounts_items.go.tmp
mv mount/model_all_of_subscription_discounts_items.go.tmp \
  mount/model_all_of_subscription_discounts_items.go

cat mount/model_all_of_subscription_discounts_items.go | \
  sed 's/MeteredComponent \*Object/MeteredComponent \MeteredComponent/g' \
  > mount/model_all_of_subscription_discounts_items.go.tmp
mv mount/model_all_of_subscription_discounts_items.go.tmp \
  mount/model_all_of_subscription_discounts_items.go

cat mount/model_all_of_customer_portal_active_subscription_subscription.go | \
  sed 's/CustomerName \*Object/CustomerName \string/g' \
  > mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp
mv mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp \
  mount/model_all_of_customer_portal_active_subscription_subscription.go

cat mount/model_all_of_customer_portal_active_subscription_subscription.go | \
  sed 's/Discounts \[]Object/Discounts []Discount/g' \
  > mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp
mv mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp \
  mount/model_all_of_customer_portal_active_subscription_subscription.go

cat mount/model_all_of_customer_portal_active_subscription_subscription.go | \
  sed 's/PricePlanName \*Object/PricePlanName \string/g' \
  > mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp
mv mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp \
  mount/model_all_of_customer_portal_active_subscription_subscription.go

cat mount/model_all_of_customer_portal_active_subscription_subscription.go | \
  sed 's/PricePlan \*Object/PricePlan \PricePlan/g' \
  > mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp
mv mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp \
  mount/model_all_of_customer_portal_active_subscription_subscription.go

cat mount/model_all_of_customer_portal_active_subscription_subscription.go | \
  sed 's/AddOns \*Object/AddOns []AddOn/g' \
  > mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp
mv mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp \
  mount/model_all_of_customer_portal_active_subscription_subscription.go

cat mount/model_all_of_customer_portal_active_subscription_subscription.go | \
  sed 's/TrialOverride \*Object/TrialOverride \Trial/g' \
  > mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp
mv mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp \
  mount/model_all_of_customer_portal_active_subscription_subscription.go

cat mount/model_all_of_customer_portal_active_subscription_subscription.go | \
  sed 's/MeteredComponent \*Object/MeteredComponent \MeteredComponent/g' \
  > mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp
mv mount/model_all_of_customer_portal_active_subscription_subscription.go.tmp \
  mount/model_all_of_customer_portal_active_subscription_subscription.go

cat mount/model_create_credit_grant_args.go | \
  sed 's/price,omitempty/price/g' \
  > mount/model_create_credit_grant_args.go.tmp
mv mount/model_create_credit_grant_args.go.tmp \
  mount/model_create_credit_grant_args.go

cat mount/model_create_credit_grant_args.go | \
  sed 's/EffectiveAt time.Time/EffectiveAt \*time.Time/g' \
  > mount/model_create_credit_grant_args.go.tmp
mv mount/model_create_credit_grant_args.go.tmp \
  mount/model_create_credit_grant_args.go

cat mount/model_create_credit_grant_args.go | \
  sed 's/ExpiresAt time.Time/ExpiresAt \*time.Time/g' \
  > mount/model_create_credit_grant_args.go.tmp
mv mount/model_create_credit_grant_args.go.tmp \
  mount/model_create_credit_grant_args.go


cat mount/model_credit_grant.go | \
  sed '/import (/,+2d' \
  > mount/model_credit_grant.go.tmp
mv mount/model_credit_grant.go.tmp \
  mount/model_credit_grant.go

#Remove unnecessary files that get generated incorrectly. Some of the Octane OpenAPI types are ambiguous
#But we don't need to bother fixing them if they aren't used in the SDK
rm mount/model_all_of_price_plan_discount.go
rm mount/model_coupon1.go
rm mount/model_all_of_subscription_discount_override.go

rm -f internal/swagger/*.go
mkdir -p internal/swagger/
cp mount/*.go internal/swagger/

