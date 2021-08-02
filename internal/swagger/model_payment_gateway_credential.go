/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PaymentGatewayCredential struct {
	// One of `STRIPE` or `PADDLE`
	PaymentGateway string `json:"payment_gateway"`
	// Payment gateway account id associated with customer
	AccountId string `json:"account_id"`
	// Unique name identifier of a customer
	AuthToken string `json:"auth_token,omitempty"`
}
