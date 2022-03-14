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

type PastInvoice struct {
	Status string `json:"status,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	ExportUrl string `json:"export_url,omitempty"`
	AmountDue float64 `json:"amount_due,omitempty"`
	DueDate time.Time `json:"due_date,omitempty"`
	StatusDescription string `json:"status_description,omitempty"`
}