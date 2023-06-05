/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateMeteredComponentLimitsInput struct {
	// Numeric limit to set on customer usage for the meter.
	Limit int32 `json:"limit,omitempty"`
	LabelLimits []MeteredComponentLabelLimitInputArgs `json:"label_limits,omitempty"`
}