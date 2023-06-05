/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CardInfo struct {
	// Month the card expires
	ExpMonth int32 `json:"exp_month,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	// Last 4 digits of the card.
	Last4 string `json:"last4,omitempty"`
	// Brand of card. E.g. Amex, Visa, etc.
	Brand string `json:"brand,omitempty"`
	// Year the card expires
	ExpYear int32 `json:"exp_year,omitempty"`
	// Country of the card
	Country string `json:"country,omitempty"`
}