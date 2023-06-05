/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateRefundArgs struct {
	// Invoice that the refund should be against
	InvoiceId int32 `json:"invoice_id,omitempty"`
	// Amount to be refunded
	Amount int32 `json:"amount,omitempty"`
	// Invoice that the refund should be against
	InvoiceUuid string `json:"invoice_uuid,omitempty"`
}
