/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MeteredComponentInputArgs struct {
	PriceScheme *PriceSchemeInputArgs `json:"price_scheme,omitempty"`
	// Numeric limit to set on customer usage for the meter.
	Limit int32 `json:"limit,omitempty"`
	MeterId int32 `json:"meter_id,omitempty"`
	// Name to be used on invoice.
	DisplayName string `json:"display_name,omitempty"`
	Id int32 `json:"id,omitempty"`
	// Codename of the meter.
	MeterName string `json:"meter_name,omitempty"`
	LabelLimits []MeteredComponentLabelLimitInputArgs `json:"label_limits,omitempty"`
}
