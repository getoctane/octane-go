/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpcomingInvoices struct {
	// The id offset to start at when paging forwards
	ForwardSecondarySortOffset int32 `json:"forward_secondary_sort_offset,omitempty"`
	SortDirection string `json:"sort_direction,omitempty"`
	// The number of items to fetch. Defaults to 10.
	Limit int32 `json:"limit,omitempty"`
	// The sort column offset to start at when paging forwards
	ForwardSortOffset string `json:"forward_sort_offset,omitempty"`
	Invoices []UpcomingInvoice `json:"invoices,omitempty"`
	SortColumn string `json:"sort_column,omitempty"`
}
