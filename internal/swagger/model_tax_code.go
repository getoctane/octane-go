/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TaxCode struct {
	// A friendly description of this tax code.
	Description string `json:"description,omitempty"`
	// The unique ID number of this tax code.
	Id string `json:"id,omitempty"`
	// The type of this tax code.
	TaxCodeTypeId string `json:"tax_code_type_id,omitempty"`
	// The Avalara Entity Use Code represented by this tax code.
	EntityUseCode string `json:"entity_use_code,omitempty"`
	// A code string that identifies this tax code.
	TaxCode string `json:"tax_code,omitempty"`
}
