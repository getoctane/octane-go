/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CustomerPortalMeter struct {
	// Name of the meter.
	MeterName string `json:"meter_name,omitempty"`
	// Type of the meter. E.g. COUNTER or GAUGE.
	MeterType string `json:"meter_type,omitempty"`
	// The raw and prettified label keys and values
	LabelsWithDisplayNames []CustomerPortalMeterLabelsWithDisplayName `json:"labels_with_display_names,omitempty"`
	// Name of the unit the meter uses.
	UnitName string `json:"unit_name,omitempty"`
	// Primary labels with keys and values
	Labels []CustomerPortalMeterLabels `json:"labels,omitempty"`
	// Display name of the meter.
	MeterDisplayName string `json:"meter_display_name,omitempty"`
}
