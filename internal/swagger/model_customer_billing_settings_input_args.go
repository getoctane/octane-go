/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CustomerBillingSettingsInputArgs struct {
	// Optional description attached to the invoice
	InvoiceMemo string `json:"invoice_memo,omitempty"`
	// Account name for ACH/Wire transfer instructions
	AchAccountName string `json:"ach_account_name,omitempty"`
	// Time length after which to attempt invoice/payment retry.
	RetryFrequencyLength int32 `json:"retry_frequency_length,omitempty"`
	// Flag that controls whether to do automated taxes via payment provider
	TaxViaPaymentProvider bool `json:"tax_via_payment_provider,omitempty"`
	// Time length unit of the grace period between the end of a billing cycle and invoice generation. Must be `day`.
	InvoiceGracePeriodUnit string `json:"invoice_grace_period_unit,omitempty"`
	// Flag that controls whether or not to invoice/charge gauge meters upfront according to their value at start of cycle. Only applies if invoice_fixed_components_at_start is enabled.
	InvoiceMeteredComponentsAtStart bool `json:"invoice_metered_components_at_start,omitempty"`
	// Flag that controls whether or not to auto-charge the customer based on the invoice.
	ChargesEnabled bool `json:"charges_enabled,omitempty"`
	// Flag that controls the number of retry attempts for invoicing/payments.
	RetryAttempts int32 `json:"retry_attempts,omitempty"`
	// Sets the due date on invoices to the number of days after the invoice is sent
	DaysUntilDue int32 `json:"days_until_due,omitempty"`
	// Default value for whether to align billing cycles to calendar on subscriptions
	AlignBillingCyclesToCalendar bool `json:"align_billing_cycles_to_calendar,omitempty"`
	// Bank name for ACH/Wire transfer instructions
	AchBankName string `json:"ach_bank_name,omitempty"`
	// Time length unit after which to attempt invoice/payment retry.
	RetryFrequencyUnit string `json:"retry_frequency_unit,omitempty"`
	// If using Stripe, this field can be used to configure whether invoices should be finalized immediately when they are created.
	StripeImmediateFinalization bool `json:"stripe_immediate_finalization,omitempty"`
	// Time length of the grace period between the end of a billing cycle and invoice generation in days.
	InvoiceGracePeriodLength int32 `json:"invoice_grace_period_length,omitempty"`
	// Swift code for ACH/Wire transfer instructions
	AchSwiftCode string `json:"ach_swift_code,omitempty"`
	// Time length unit of the grace period between the end of invoice generation and actual charge. One of `minute`, `hour`, `day`.
	PaymentGracePeriodUnit string `json:"payment_grace_period_unit,omitempty"`
	// The percentage tax rate to apply to invoices.
	TaxRate float64 `json:"tax_rate,omitempty"`
	// Account number for ACH/Wire transfer instructions
	AchAccountNumber string `json:"ach_account_number,omitempty"`
	// Time length of the grace period between the end of invoice generation and the actual charge. *NOTE*: The specified length is unitless. Unit is designated with the `payment_grace_period_unit` field.
	PaymentGracePeriodLength int32 `json:"payment_grace_period_length,omitempty"`
	// Flag that controls whether invoices are auto-approved or require manual approval
	AutoApproveInvoices bool `json:"auto_approve_invoices,omitempty"`
	// First line of bank address for ACH/Wire transfer instructions
	AchBankAddress1 string `json:"ach_bank_address_1,omitempty"`
	// Flag that controls whether or not to invoice/charge the base rate, add ons and other fixed price plan components at the beginning of the billing cycle.
	InvoiceFixedComponentsAtStart bool `json:"invoice_fixed_components_at_start,omitempty"`
	// Second line of bank address for ACH/Wire transfer instructions
	AchBankAddress2 string `json:"ach_bank_address_2,omitempty"`
	// Flag that controls whether or not invoices should be sent to customers.
	ShouldSendInvoiceToCustomers bool `json:"should_send_invoice_to_customers,omitempty"`
	// Flag determining whether ACH/Wire instructions should be included on invoices.
	IncludeAchInstructions bool `json:"include_ach_instructions,omitempty"`
	// If using stripe, this field can be used to configure whether invoices should be auto advanced for collection
	StripeAutoAdvance bool `json:"stripe_auto_advance,omitempty"`
	// True if customer updates should be synced to Stripe.
	SyncCustomerDataToPaymentGateway bool `json:"sync_customer_data_to_payment_gateway,omitempty"`
	// Flag that controls whether to invoice through Octane or through payment provider
	InvoiceViaOctane bool `json:"invoice_via_octane,omitempty"`
	CustomerInvoiceDetailLevel string `json:"customer_invoice_detail_level,omitempty"`
	// Flag that controls whether or not to invoice/charge a true up for a billing cycle on the following invoice. Only applies if invoice_fixed_components_at_start is enabled.
	InvoiceOverages bool `json:"invoice_overages,omitempty"`
	// Optional url of a custom image to include on invoices.
	InvoiceLogoUrl string `json:"invoice_logo_url,omitempty"`
	// ABA/Routing number for ACH/Wire transfer instructions
	AchRoutingNumber string `json:"ach_routing_number,omitempty"`
}
