/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateSelfServeSettingsArgs struct {
	// Price per credit, in cents, that the customer is charged for buying credits through the customer portal
	PricePerCreditCents int32 `json:"price_per_credit_cents,omitempty"`
	// Time length unit for the default expiration for credits bought in the customer portal.
	CreditsExpirationUnit string `json:"credits_expiration_unit,omitempty"`
	// True if the customer can purchase credits via self serve. Defaults to False.
	PurchaseCredits bool `json:"purchase_credits,omitempty"`
	// True if the vendor has enabled customization for their customer portal.
	Customization bool `json:"customization,omitempty"`
	// Time length of the default expiration for credits bought in the customer portal.
	CreditsExpirationLength int32 `json:"credits_expiration_length,omitempty"`
	// True if the customer can switch their current price plan via self serve. Defaults to False.
	SwitchPricePlans bool `json:"switch_price_plans,omitempty"`
	// True if the vendor has enabled customization for their customer portal.
	Enabled bool `json:"enabled,omitempty"`
}
