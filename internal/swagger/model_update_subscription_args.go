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
	EffectiveAt time.Time `json:"effective_at,omitempty"`
	PricePlanTag string `json:"price_plan_tag,omitempty"`
	VendorId int32 `json:"vendor_id,omitempty"`
	// Boolean that indicates whether to update the subscription at the end of thebilling cycle. If 'true' and `effective_at` is set, will return an error.
	AtCycleEnd bool `json:"at_cycle_end,omitempty"`
	LimitsOverride []LimitInputArgs `json:"limits_override,omitempty"`
	AlignToCalendar bool `json:"align_to_calendar,omitempty"`
	DiscountOverride *AllOfUpdateSubscriptionArgsDiscountOverride `json:"discount_override,omitempty"`
	TrialOverride *AllOfUpdateSubscriptionArgsTrialOverride `json:"trial_override,omitempty"`
	CouponOverrideId int32 `json:"coupon_override_id,omitempty"`
	PricePlanId int32 `json:"price_plan_id,omitempty"`
	CouponOverrideName string `json:"coupon_override_name,omitempty"`
	AddOns []SubscriptionAddOnInput `json:"add_ons,omitempty"`
	FeaturesOverride []FeatureInputArgs `json:"features_override,omitempty"`
	PricePlanName string `json:"price_plan_name,omitempty"`
	CustomerId int32 `json:"customer_id,omitempty"`
}
