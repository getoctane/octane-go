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
	Country string `json:"country,omitempty"`
	AddressLine1 string `json:"address_line_1,omitempty"`
	State string `json:"state,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
	Email string `json:"email,omitempty"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	City string `json:"city,omitempty"`
	Url string `json:"url,omitempty"`
	Phone string `json:"phone,omitempty"`
}
