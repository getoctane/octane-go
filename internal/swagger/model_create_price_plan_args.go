/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreatePricePlanArgs struct {
	Trial *TrialInputArgs `json:"trial,omitempty"`
	// The frequency (as a an integer multiple of the period) at which to charge the minimum charge.
	MinimumChargeFrequency int32 `json:"minimum_charge_frequency,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Description string `json:"description,omitempty"`
	AddOns []AddOnInputArgs `json:"add_ons,omitempty"`
	VendorId int32 `json:"vendor_id,omitempty"`
	// Custom invoice description for the base price line item.
	BasePriceDescription string `json:"base_price_description,omitempty"`
	Features []FeatureInputArgs `json:"features,omitempty"`
	// The frequency (as a an integer multiple of the period) at which to charge the base price.
	BasePriceFrequency int32 `json:"base_price_frequency,omitempty"`
	Name string `json:"name,omitempty"`
	BasePrice int32 `json:"base_price,omitempty"`
	Period string `json:"period,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	MeteredComponents []MeteredComponentInputArgs `json:"metered_components,omitempty"`
	Limits []LimitInputArgs `json:"limits,omitempty"`
	// Minimum amount (in cents) to charge every price plan period.
	MinimumCharge int32 `json:"minimum_charge,omitempty"`
}
