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

type UpcomingInvoice struct {
	Status string `json:"status,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	GenerateDate time.Time `json:"generate_date,omitempty"`
	ExportUrl string `json:"export_url,omitempty"`
	StatusDescription string `json:"status_description,omitempty"`
}