/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateBillingSettingsInputArgs struct {
	PaymentGracePeriodLength int32 `json:"payment_grace_period_length,omitempty"`
	InvoiceGracePeriodUnit string `json:"invoice_grace_period_unit,omitempty"`
	RetryFrequencyLength int32 `json:"retry_frequency_length,omitempty"`
	CustomerInvoiceDetailLevel string `json:"customer_invoice_detail_level,omitempty"`
	InvoiceFixedComponentsAtStart bool `json:"invoice_fixed_components_at_start,omitempty"`
	RetryAttempts int32 `json:"retry_attempts,omitempty"`
	TaxRate float64 `json:"tax_rate,omitempty"`
	InvoiceGracePeriodLength int32 `json:"invoice_grace_period_length,omitempty"`
	ChargesEnabled bool `json:"charges_enabled,omitempty"`
	CustomerId int32 `json:"customer_id,omitempty"`
	AutoApproveInvoices bool `json:"auto_approve_invoices,omitempty"`
	VendorId int32 `json:"vendor_id,omitempty"`
	InvoiceViaOctane bool `json:"invoice_via_octane,omitempty"`
	ShouldSendInvoiceToCustomers bool `json:"should_send_invoice_to_customers,omitempty"`
	RetryFrequencyUnit string `json:"retry_frequency_unit,omitempty"`
	PaymentGracePeriodUnit string `json:"payment_grace_period_unit,omitempty"`
}
