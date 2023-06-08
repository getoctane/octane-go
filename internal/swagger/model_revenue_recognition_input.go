/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RevenueRecognitionInput struct {
	// List of customer names for which to compute booked/recognized revenue.
	CustomerNames []string `json:"customer_names,omitempty"`
}
