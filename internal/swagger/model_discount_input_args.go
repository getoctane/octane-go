/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DiscountInputArgs struct {
	DiscountType string `json:"discount_type,omitempty"`
	// For ADD_ON scoped discounts: the name of the add on that the discount covers.
	AddOnName string `json:"add_on_name,omitempty"`
	// Length, in billing cycles, that this discount will be active.
	BillingCycleDuration int32 `json:"billing_cycle_duration,omitempty"`
	// For METERED_COMPONENT scoped discounts: Dictionary of labels (key: value) that the discount covers. The entire set of labels must be provided.
	Labels map[string]string `json:"labels,omitempty"`
	// The scope that this discount covers. One of 'INVOICE_TOTAL', 'ADD_ON', 'METERED_COMPONENT'.
	Scope string `json:"scope,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	// Offset, in number of billing cycles, for when this discount will apply. If set to 0, the discount will start applying from the current billing cycle. If set to 1, the discount will start applying from the next billing cycle, etc. For scheduled subscriptions, the offset starts from the initial billing cycle.
	BillingCycleStartOffset int32 `json:"billing_cycle_start_offset,omitempty"`
	// For METERED_COMPONENT scoped discounts: the UUID of the metered component that the discount covers.
	MeteredComponentUuid string `json:"metered_component_uuid,omitempty"`
}
