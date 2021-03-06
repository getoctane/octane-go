/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type UpdateSubscriptionArgs struct {
	TrialOverride *AllOfUpdateSubscriptionArgsTrialOverride `json:"trial_override,omitempty"`
	VendorId int32 `json:"vendor_id,omitempty"`
	AddOns []SubscriptionAddOnInput `json:"add_ons,omitempty"`
	CouponOverrideName string `json:"coupon_override_name,omitempty"`
	EffectiveAt time.Time `json:"effective_at,omitempty"`
	AlignToCalendar bool `json:"align_to_calendar,omitempty"`
	DiscountOverride *AllOfUpdateSubscriptionArgsDiscountOverride `json:"discount_override,omitempty"`
	FeaturesOverride []FeatureInputArgs `json:"features_override,omitempty"`
	CouponOverrideId int32 `json:"coupon_override_id,omitempty"`
	PricePlanName string `json:"price_plan_name,omitempty"`
	// Boolean that indicates whether to update the subscription at the end of thebilling cycle. If 'true' and `effective_at` is set, will return an error.
	AtCycleEnd bool `json:"at_cycle_end,omitempty"`
	PricePlanId int32 `json:"price_plan_id,omitempty"`
	PricePlanTag string `json:"price_plan_tag,omitempty"`
	CustomerId int32 `json:"customer_id,omitempty"`
	LimitsOverride []LimitInputArgs `json:"limits_override,omitempty"`
}
