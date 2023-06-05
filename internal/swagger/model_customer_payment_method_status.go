/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CustomerPaymentMethodStatus struct {
	// The status of the customer's current payment method. Can be one of: NO_PAYMENT_GATEWAY_CREDENTIAL, NO_PAYMENT_METHOD, EXPIRED_PAYMENT_METHOD, VALID_PAYMENT_METHOD
	Status string `json:"status,omitempty"`
}
