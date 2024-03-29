/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type CreateCustomerArgs struct {
	PricePlanName string `json:"price_plan_name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	ContactInfo *ContactInfoInputArgs `json:"contact_info,omitempty"`
	PricePlanTag string `json:"price_plan_tag,omitempty"`
	VendorId int32 `json:"vendor_id,omitempty"`
	AutogenerateAccountingCustomer bool `json:"autogenerate_accounting_customer,omitempty"`
	MeasurementMappings []CustomerMeasurementMappingInputArgs `json:"measurement_mappings,omitempty"`
	Name string `json:"name,omitempty"`
	AutogeneratePaymentGatewayCustomer bool `json:"autogenerate_payment_gateway_customer,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	CustomerMetadata []CustomerMetadataInput `json:"customer_metadata,omitempty"`
}
