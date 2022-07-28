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
	VendorId int32 `json:"vendor_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	BasePrice int32 `json:"base_price,omitempty"`
	MeteredComponents []MeteredComponentInputArgs `json:"metered_components,omitempty"`
	Discount *DiscountInputArgs `json:"discount,omitempty"`
	Trial *TrialInputArgs `json:"trial,omitempty"`
	Limits []LimitInputArgs `json:"limits,omitempty"`
	Period string `json:"period,omitempty"`
	Name string `json:"name,omitempty"`
	AddOns []AddOnInputArgs `json:"add_ons,omitempty"`
	Features []FeatureInputArgs `json:"features,omitempty"`
	Description string `json:"description,omitempty"`
	CouponName string `json:"coupon_name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	// The frequency (as a an integer multiple of the period) at which to charge the base price.
	BasePriceFrequency int32 `json:"base_price_frequency,omitempty"`
}
