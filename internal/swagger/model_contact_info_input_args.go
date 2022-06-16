/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ContactInfoInputArgs struct {
	Phone string `json:"phone,omitempty"`
	Url string `json:"url,omitempty"`
	LegalName string `json:"legal_name,omitempty"`
	City string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
	AddressLine1 string `json:"address_line_1,omitempty"`
	Email string `json:"email,omitempty"`
	LogoUrl string `json:"logo_url,omitempty"`
	State string `json:"state,omitempty"`
	AddressLine2 string `json:"address_line_2,omitempty"`
}
