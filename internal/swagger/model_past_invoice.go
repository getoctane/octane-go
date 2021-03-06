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
	AmountDue float64 `json:"amount_due,omitempty"`
	StatusDescription string `json:"status_description,omitempty"`
	Status string `json:"status,omitempty"`
	ExportUrl string `json:"export_url,omitempty"`
	DueDate time.Time `json:"due_date,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
}
