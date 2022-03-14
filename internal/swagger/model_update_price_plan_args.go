/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdatePricePlanArgs struct {
	BasePrice int32 `json:"base_price,omitempty"`
	Description string `json:"description,omitempty"`
	Features []FeatureInputArgs `json:"features,omitempty"`
	Trial *TrialInputArgs `json:"trial,omitempty"`
	Discount *DiscountInputArgs `json:"discount,omitempty"`
	Tags []string `json:"tags,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	CouponName string `json:"coupon_name,omitempty"`
	Limits []LimitInputArgs `json:"limits,omitempty"`
	Name string `json:"name,omitempty"`
	MeteredComponents []MeteredComponentInputArgs `json:"metered_components,omitempty"`
	AddOns []AddOnInputArgs `json:"add_ons,omitempty"`
	Period string `json:"period,omitempty"`
	VendorId int32 `json:"vendor_id,omitempty"`
}
