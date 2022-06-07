/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SubscriptionAddOnInput struct {
	Quantity int32 `json:"quantity,omitempty"`
	Name string `json:"name"`
	// Override for the add-on price on this subscription.
	Price float64 `json:"price,omitempty"`
}
