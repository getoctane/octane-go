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

// List of discounts associated with this subscription.
type AllOfActiveSubscriptionDiscountsItems struct {
	DiscountType string `json:"discount_type,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	// The date when the discount is applied from.
	StartDate time.Time `json:"start_date,omitempty"`
	// The date when the discount ends.
	EndDate time.Time `json:"end_date,omitempty"`
	// Offset in number of billing cycles for when this discount will apply. For example, if set to 1, the discount will apply from the start of the next billing cycle.
	BillingCycleStartOffset int32 `json:"billing_cycle_start_offset,omitempty"`
	// Duration of this discount in number of billing cycles.
	BillingCycleDuration int32 `json:"billing_cycle_duration,omitempty"`
	// The id of coupon associated with this discount, none if discount does not originate from coupon
	CouponId int32 `json:"coupon_id,omitempty"`
	// External facing unique identifier of a price plan
	ExternalUuid string `json:"external_uuid,omitempty"`
	// The scope that this discount covers. One of 'INVOICE_TOTAL', 'ADD_ON', 'METERED_COMPONENT'.
	Scope string `json:"scope"`
	// Add-on this discount covers if scope is ADD_ON.
	AddOn AddOn `json:"add_on,omitempty"`
	// Metered Component this discount covers if scope is METERED_COMPONENT.
	MeteredComponent MeteredComponent `json:"metered_component,omitempty"`
	// Dictionary of labels (key: value) to which the discount applies if scope is METERED_COMPONENT.
	Labels map[string]string `json:"labels,omitempty"`
}
