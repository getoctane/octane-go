/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CustomerLabelLimit struct {
	Limit float64 `json:"limit,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}