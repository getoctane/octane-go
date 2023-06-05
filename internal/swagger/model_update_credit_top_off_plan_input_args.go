/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateCreditTopOffPlanInputArgs struct {
	// Time length of the default expiration for credits granted in a top off.
	ExpirationLength int32 `json:"expiration_length,omitempty"`
	// Price for the grant, in lowest denomination (i.e cents).
	Price int32 `json:"price,omitempty"`
	// The threshold in amount of credits at which the balance will be topped off.
	TriggerAmount int32 `json:"trigger_amount,omitempty"`
	// Time length unit for the default expiration for credits granted in a top off.
	ExpirationUnit string `json:"expiration_unit,omitempty"`
	// A description that will be used on the invoice line items.
	Description string `json:"description,omitempty"`
	// Amount of credits that are granted in a single top off.
	GrantAmount int32 `json:"grant_amount,omitempty"`
}