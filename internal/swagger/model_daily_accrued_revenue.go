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

type DailyAccruedRevenue struct {
	// The start time of the billing cycle for which the accrued revenue is computed
	StartTime time.Time `json:"start_time,omitempty"`
	LineItems []AccruedRevenueLineItem `json:"line_items,omitempty"`
	// The end time till when the accured revenue is computed
	EndTime time.Time `json:"end_time,omitempty"`
	// The date for which this accrued revenue is computed
	Date time.Time `json:"date,omitempty"`
	// Total accrued revenue for the day in cents
	TotalAccruedRevenue int32 `json:"total_accrued_revenue,omitempty"`
}