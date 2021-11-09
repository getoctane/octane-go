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

type Invoice struct {
	// Amount due before any credits are applied
	SubTotal int32 `json:"sub_total,omitempty"`
	Id string `json:"id,omitempty"`
	LatestInvoiceAttemptAt time.Time `json:"latest_invoice_attempt_at,omitempty"`
	// Non-empty string if there was an error while processing payment
	PaymentError string `json:"payment_error,omitempty"`
	// The number of retries done to send the invoice
	InvoiceRetryAttempt int32 `json:"invoice_retry_attempt,omitempty"`
	LineItems []LineItems `json:"line_items,omitempty"`
	// False if not paid yet
	IsPaid bool `json:"is_paid,omitempty"`
	// False if invoice has not been sent to the customer
	IsInvoiced bool `json:"is_invoiced,omitempty"`
	// Total amount due
	AmountDue int32 `json:"amount_due,omitempty"`
	LatestPaymentAttemptAt time.Time `json:"latest_payment_attempt_at,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	DueDate time.Time `json:"due_date,omitempty"`
	// Non-empty string if there was an error while sending out invoice
	InvoicingError string `json:"invoicing_error,omitempty"`
	// False if not approved
	IsApproved bool `json:"is_approved,omitempty"`
	// Any discount credits applied to the invoice
	DiscountCredit int32 `json:"discount_credit,omitempty"`
	EndTime time.Time `json:"end_time,omitempty"`
	// The number of retries done to process the payment
	PaymentRetryAttempt int32 `json:"payment_retry_attempt,omitempty"`
}
