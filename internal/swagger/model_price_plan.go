/*
 * Octane API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type PricePlan struct {
	// External facing unique identifier of a price plan
	ExternalUuid string `json:"external_uuid,omitempty"`
	// Unique name indentifier of a price plan
	Name string `json:"name"`
	// UI-friendly name used for data display. Defaults to `name`.
	DisplayName string `json:"display_name,omitempty"`
	Description string `json:"description,omitempty"`
	// Lowest denomination of currency. e.g. USD is represented as cents.
	BasePrice float64 `json:"base_price,omitempty"`
	BasePriceFrequency int32 `json:"base_price_frequency,omitempty"`
	// Custom invoice description for the base price line item.
	BasePriceDescription string `json:"base_price_description,omitempty"`
	// Time period that defines the length of a price plan cycle. One of `day`, `week`, `month`, `quarter`, or `year`.
	Period string `json:"period"`
	MeteredComponents []MeteredComponent `json:"metered_components"`
	// Minimum amount to charge every 'period'
	MinimumCharge int32 `json:"minimum_charge,omitempty"`
	// The frequency (as a an integer multiple of the period) at which to charge the minimum charge.
	MinimumChargeFrequency int32 `json:"minimum_charge_frequency,omitempty"`
	Features []Feature `json:"features,omitempty"`
	AddOns []AddOn `json:"add_ons,omitempty"`
	Limits []Limit `json:"limits,omitempty"`
	Tags []PricePlanTag `json:"tags,omitempty"`
	Trial *Trial `json:"trial,omitempty"`
	// ISO-8601 formatted creation timestamp of price plan version
	CreatedAt time.Time `json:"created_at,omitempty"`
}
