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

// Usage by day.
type AllOfCycleUsageUsageByTimeItems struct {
	// Total usage during this window.
	Usage int32 `json:"usage,omitempty"`
	// Label value. Only present if label_group_by is provided.
	LabelValue string `json:"label_value,omitempty"`
	// Start of the 24 hour time window in UTC.
	Time time.Time `json:"time,omitempty"`
	// Label key. Only present if label_group_by is provided.
	LabelKey string `json:"label_key,omitempty"`
}
