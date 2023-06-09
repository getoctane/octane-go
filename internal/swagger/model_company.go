/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Company struct {
	// This flag is true if this company is the default company for this account.
	IsDefault bool `json:"is_default,omitempty"`
	// The unique ID number of this company.
	Id string `json:"id,omitempty"`
	// A unique code that references this company within your account.
	CompanyCode string `json:"company_code,omitempty"`
	// This flag indicates whether tax activity can occur for this company.
	IsActive bool `json:"is_active,omitempty"`
	// The name of this company, as shown to customers.
	CompanyName string `json:"company_name,omitempty"`
}
