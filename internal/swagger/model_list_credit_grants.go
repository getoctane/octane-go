/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ListCreditGrants struct {
	// The number of items to fetch. Defaults to 10.
	Limit int32 `json:"limit,omitempty"`
	CreditGrants []CreditGrant `json:"credit_grants,omitempty"`
	SortDirection string `json:"sort_direction,omitempty"`
	// The unique offset to start at when paging forwards
	ForwardSecondarySortOffset string `json:"forward_secondary_sort_offset,omitempty"`
	SortColumn string `json:"sort_column,omitempty"`
	// The sort column offset to start at when paging forwards
	ForwardSortOffset string `json:"forward_sort_offset,omitempty"`
}
