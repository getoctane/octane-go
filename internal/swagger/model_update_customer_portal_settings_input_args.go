/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateCustomerPortalSettingsInputArgs struct {
	VendorId int32 `json:"vendor_id,omitempty"`
	PricePlanTagsFilter string `json:"price_plan_tags_filter,omitempty"`
	PricePlanNamesFilter string `json:"price_plan_names_filter,omitempty"`
}
