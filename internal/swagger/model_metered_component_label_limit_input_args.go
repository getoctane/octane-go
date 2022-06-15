/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MeteredComponentLabelLimitInputArgs struct {
	// Numeric limit to set on customer usage for the meter with the given labels.
	Limit float64 `json:"limit"`
	// Dictionary of labels (key: value) to which the limit applies. A value of 'any' will apply the limit to any single value of the field.
	Labels map[string]string `json:"labels"`
}