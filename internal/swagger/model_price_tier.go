/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PriceTier struct {
	// The price (in lowest currency denomination by which to charge, given that the usage is within the cap range.
	Price float64 `json:"price"`
	// Cap of the tier, meaning that any subsequent usage will be bucketed into the following tier. If cap is undefined, it is effectively treated as Infinity.
	Cap float64 `json:"cap,omitempty"`
	// The line item description to use if usage falls in this tier.
	Description string `json:"description,omitempty"`
}
