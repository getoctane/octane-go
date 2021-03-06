/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateBillingSettingsInputArgs struct {
	VendorId int32 `json:"vendor_id,omitempty"`
	StripeAutoAdvance bool `json:"stripe_auto_advance,omitempty"`
	ChargesEnabled bool `json:"charges_enabled,omitempty"`
	InvoiceFixedComponentsAtStart bool `json:"invoice_fixed_components_at_start,omitempty"`
	InvoiceMemo string `json:"invoice_memo,omitempty"`
	CustomerInvoiceDetailLevel string `json:"customer_invoice_detail_level,omitempty"`
	RetryFrequencyLength int32 `json:"retry_frequency_length,omitempty"`
	RetryAttempts int32 `json:"retry_attempts,omitempty"`
	TaxRate float64 `json:"tax_rate,omitempty"`
	DaysUntilDue int32 `json:"days_until_due,omitempty"`
	PaymentGracePeriodLength int32 `json:"payment_grace_period_length,omitempty"`
	InvoiceOverages bool `json:"invoice_overages,omitempty"`
	AutoApproveInvoices bool `json:"auto_approve_invoices,omitempty"`
	ShouldSendInvoiceToCustomers bool `json:"should_send_invoice_to_customers,omitempty"`
	CustomerId int32 `json:"customer_id,omitempty"`
	InvoiceMeteredComponentsAtStart bool `json:"invoice_metered_components_at_start,omitempty"`
	PaymentGracePeriodUnit string `json:"payment_grace_period_unit,omitempty"`
	InvoiceGracePeriodUnit string `json:"invoice_grace_period_unit,omitempty"`
	InvoiceGracePeriodLength int32 `json:"invoice_grace_period_length,omitempty"`
	InvoiceViaOctane bool `json:"invoice_via_octane,omitempty"`
	RetryFrequencyUnit string `json:"retry_frequency_unit,omitempty"`
	TaxViaPaymentProvider bool `json:"tax_via_payment_provider,omitempty"`
}
