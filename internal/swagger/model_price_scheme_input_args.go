/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PriceSchemeInputArgs struct {
	// One of 'FLAT', 'VOLUME', or 'STAIRSTEP'
	SchemeType string `json:"scheme_type,omitempty"`
	// The name of the unit used for this metered component (e.g., gigabyte)
	UnitName string `json:"unit_name,omitempty"`
	// Array of (key, value) meter labels to price on & the price tiers that should be used against those labels
	PriceList []interface{} `json:"price_list,omitempty"`
	// The time unit for the metered component (e.g., month or hour)
	TimeUnitName string `json:"time_unit_name,omitempty"`
	// Array of price tiers, each of which consists of `price` and `cap` key:value pairs
	Prices []PriceInputArgs `json:"prices,omitempty"`
}
