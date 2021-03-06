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

type Discount struct {
	DiscountType string `json:"discount_type,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	// The date when the discount is applied from.
	StartDate time.Time `json:"start_date,omitempty"`
	// The date when the discount ends.
	EndDate time.Time `json:"end_date,omitempty"`
	// The id of coupon associated with this discount, none if discount does not originate from coupon
	CouponId int32 `json:"coupon_id,omitempty"`
}
