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

type DeleteSubscriptionArgs struct {
	VendorId int32 `json:"vendor_id,omitempty"`
	ExpireAt time.Time `json:"expire_at,omitempty"`
	CustomerId int32 `json:"customer_id,omitempty"`
	// Boolean that indicates whether to expire the subscription at the end of thebilling cycle. If 'true' and `expire_at` is set, will return an error.
	AtCycleEnd bool `json:"at_cycle_end,omitempty"`
}
