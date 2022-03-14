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
	MeterId int32 `json:"meter_id,omitempty"`
	Id int32 `json:"id,omitempty"`
	MeterName string `json:"meter_name,omitempty"`
	PriceScheme *PriceSchemeInputArgs `json:"price_scheme,omitempty"`
	Limit int32 `json:"limit,omitempty"`
}
