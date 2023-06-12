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

type SubscriptionAddOnItem struct {
	// Boolean that indicates whether to update the subscription add on at the start of the billing cycle. If 'true' and either of `effective_at` or `at_cycle_end` are set, will return an error.
	AtCycleStart bool `json:"at_cycle_start,omitempty"`
	// Boolean that indicates whether to update the subscription add on at the end of the billing cycle. If 'true' and either of `effective_at` or `at_cycle_start` are set, will return an error.
	AtCycleEnd bool `json:"at_cycle_end,omitempty"`
	// Quantity represents how many of this add on you want to attach to the subscription. Can be positive forincreasing the number of this add on or negative for decreasing.
	Quantity int32 `json:"quantity,omitempty"`
	EffectiveAt time.Time `json:"effective_at,omitempty"`
	FeatureName string `json:"feature_name,omitempty"`
}
