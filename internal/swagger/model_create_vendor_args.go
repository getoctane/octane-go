/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateVendorArgs struct {
	VendorDisplayName string `json:"vendor_display_name,omitempty"`
	VendorName string `json:"vendor_name,omitempty"`
	ApiKey string `json:"api_key,omitempty"`
	Name string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	ContactInfo *ContactInfoInputArgs `json:"contact_info,omitempty"`
}
