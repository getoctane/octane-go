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

type Subscription struct {
	Uuid string `json:"uuid"`
	// Unique name identifier of a customer
	CustomerName string `json:"customer_name"`
	// Unique name indentifier of a price plan
	PricePlanName string `json:"price_plan_name"`
	// Price plan associated with this subscription.
	PricePlan *AllOfSubscriptionPricePlan `json:"price_plan,omitempty"`
	// Align billing cycles to a calendar unit if true. For example if the period is month, cycles will end on the first of every month.
	AlignToCalendar bool `json:"align_to_calendar,omitempty"`
	Discounts []AllOfSubscriptionDiscountsItems `json:"discounts,omitempty"`
	AddOns AddOn `json:"add_ons,omitempty"`
	// Optional trial override for the associated subscription.
	TrialOverride *AllOfSubscriptionTrialOverride `json:"trial_override,omitempty"`
	// Optional base price override for the associated subscription.
	BasePriceOverride float64 `json:"base_price_override,omitempty"`
	FeaturesOverride []Feature `json:"features_override,omitempty"`
	LimitsOverride []Limit `json:"limits_override,omitempty"`
	// ISO-8601 formatted timestamp that defines when the subscription should take effect. If this field is omitted, the subscription is effective upon creation.
	EffectiveAt time.Time `json:"effective_at,omitempty"`
	// ISO-8601 formatted timestamp that defines when the subscription will expire.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}
