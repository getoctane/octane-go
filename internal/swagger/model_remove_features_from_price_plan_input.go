/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RemoveFeaturesFromPricePlanInput struct {
	// List of feature names to remove
	FeatureNames []string `json:"feature_names"`
}
