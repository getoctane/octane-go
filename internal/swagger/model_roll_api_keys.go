/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RollApiKeys struct {
	Success bool `json:"success,omitempty"`
	// The newly generated API Key.
	ApiKey string `json:"api_key,omitempty"`
}
