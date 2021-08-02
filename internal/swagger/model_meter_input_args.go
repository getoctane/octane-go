/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MeterInputArgs struct {
	VendorId int32 `json:"vendor_id,omitempty"`
	Description string `json:"description,omitempty"`
	PrimaryLabels []string `json:"primary_labels,omitempty"`
	MeterType string `json:"meter_type,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	IsIncremental bool `json:"is_incremental,omitempty"`
	Name string `json:"name,omitempty"`
	UnitName string `json:"unit_name,omitempty"`
	ExpectedLabels []string `json:"expected_labels,omitempty"`
}
