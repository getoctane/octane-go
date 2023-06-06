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

// Daily usage across the previous billing cycle.
type AllOfCustomerPortalUsagePreviousCycleUsage struct {
	UsageByTime []CycleUsage `json:"usage_by_time,omitempty"`
	// The start of the billing cycle in UTC.
	CycleStart time.Time `json:"cycle_start,omitempty"`
	// The end of the billing cycle in UTC.
	CycleEnd time.Time `json:"cycle_end,omitempty"`
	// Total usage in the cycle.
	TotalUsage int32 `json:"total_usage,omitempty"`
}
