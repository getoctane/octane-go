/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PaymentGatewayCredentialInputArgs struct {
	PaymentGateway string `json:"payment_gateway,omitempty"`
	AuthToken string `json:"auth_token,omitempty"`
	AccountId string `json:"account_id,omitempty"`
}
