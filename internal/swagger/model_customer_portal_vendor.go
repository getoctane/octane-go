/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CustomerPortalVendor struct {
	// Currency preference of the Vendor.
	Currency string `json:"currency,omitempty"`
	// Full contact info for the Vendor
	ContactInfo *AllOfCustomerPortalVendorContactInfo `json:"contact_info,omitempty"`
	// Vendor's current payment gateway.
	PaymentGateway string `json:"payment_gateway,omitempty"`
	// Unique name identifier of a Vendor
	Name string `json:"name,omitempty"`
	// Display name for the Vendor
	DisplayName string `json:"display_name,omitempty"`
}
