/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PriceScheme struct {
	DisplayName string `json:"display_name,omitempty"`
	Name string `json:"name,omitempty"`
	SchemeType string `json:"scheme_type,omitempty"`
	// Array of price tiers, each of which consists of `price` and `cap` key:value pairs
	Prices []PriceTier `json:"prices,omitempty"`
	TimeUnitName string `json:"time_unit_name,omitempty"`
	UnitName string `json:"unit_name,omitempty"`
}
