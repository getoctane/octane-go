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
	ExpectedLabels []string `json:"expected_labels,omitempty"`
	VendorId int32 `json:"vendor_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	UnitName string `json:"unit_name,omitempty"`
	IsIncremental bool `json:"is_incremental,omitempty"`
	Name string `json:"name,omitempty"`
	PrimaryLabels []string `json:"primary_labels,omitempty"`
	Description string `json:"description,omitempty"`
	MeterType string `json:"meter_type,omitempty"`
}
