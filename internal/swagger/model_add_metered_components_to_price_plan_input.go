/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AddMeteredComponentsToPricePlanInput struct {
	// List of metered components to add
	MeteredComponents []MeteredComponentInputArgs `json:"metered_components"`
}
