/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Full contact info for the Vendor
type AllOfCustomerPortalVendorContactInfo struct {
	AddressLine1 string `json:"address_line_1,omitempty"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
	Url string `json:"url,omitempty"`
	LogoUrl string `json:"logo_url,omitempty"`
	Email string `json:"email,omitempty"`
	SecondaryEmails string `json:"secondary_emails,omitempty"`
	Phone string `json:"phone,omitempty"`
	LegalName string `json:"legal_name,omitempty"`
	VatId string `json:"vat_id,omitempty"`
}