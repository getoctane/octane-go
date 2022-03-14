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

type ActiveSubscription struct {
	// Unique name identifier of a customer
	CustomerName string `json:"customer_name"`
	// Unique name indentifier of a price plan
	PricePlanName string `json:"price_plan_name"`
	// Price plan associated with this subscription.
	PricePlan *AllOfActiveSubscriptionPricePlan `json:"price_plan,omitempty"`
	// Optional discount override for the associated subscription.
	DiscountOverride *AllOfActiveSubscriptionDiscountOverride `json:"discount_override,omitempty"`
	AddOns []SubscriptionAddOn `json:"add_ons,omitempty"`
	// Optional trial override for the associated subscription.
	TrialOverride *AllOfActiveSubscriptionTrialOverride `json:"trial_override,omitempty"`
	FeaturesOverride []Feature `json:"features_override,omitempty"`
	LimitsOverride []Limit `json:"limits_override,omitempty"`
	// ISO-8601 formatted timestamp that defines when the subscription should take effect. If this field is omitted, the subscription is effective upon creation.
	EffectiveAt time.Time `json:"effective_at,omitempty"`
	// Align billing cycles to a calendar unit if true. For example if the period is month, cycles will end on the first of every month.
	AlignToCalendar bool `json:"align_to_calendar,omitempty"`
	// Optional base price override for the associated subscription.
	BasePriceOverride float64 `json:"base_price_override,omitempty"`
	CurrentBillingCycle *BillingCycleDate `json:"current_billing_cycle,omitempty"`
}
