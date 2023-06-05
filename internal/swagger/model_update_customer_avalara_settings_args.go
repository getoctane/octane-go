/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateCustomerAvalaraSettingsArgs struct {
	// Entity code describing this customer.
	EntityUseCode string `json:"entity_use_code,omitempty"`
	// Tax exemption number specific to this customer
	ExemptionNumber string `json:"exemption_number,omitempty"`
	// True if Avalara integration should be enabled for this customer, False otherwise.
	EnableIntegration bool `json:"enable_integration,omitempty"`
}