/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApplyCouponInputArgs struct {
	CustomerName string `json:"customer_name,omitempty"`
	VendorId int32 `json:"vendor_id,omitempty"`
	Name string `json:"name,omitempty"`
	CustomerId int32 `json:"customer_id,omitempty"`
	Code string `json:"code,omitempty"`
}
