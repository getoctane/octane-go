/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModelError struct {
	// Error message
	Message string `json:"message,omitempty"`
	// Error code
	Code int32 `json:"code,omitempty"`
	// Errors
	Errors *interface{} `json:"errors,omitempty"`
	// Error name
	Status string `json:"status,omitempty"`
}
