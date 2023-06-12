/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EntityUseCode struct {
	// Text describing the meaning of this use code.
	Description string `json:"description,omitempty"`
	// The Avalara-recognized entity use code for this definition.
	Code string `json:"code,omitempty"`
	// The name of this entity use code.
	Name string `json:"name,omitempty"`
	// A list of countries where this use code is valid.
	ValidCountries []string `json:"valid_countries,omitempty"`
}
