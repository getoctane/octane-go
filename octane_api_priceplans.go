package octane

import (
	"context"
	"net/http"

	"github.com/getoctane/octane-go/internal/swagger"
)

type (
	AddOnInputArgs struct {
		Feature *FeatureInputArgs `json:"feature,omitempty"`
		Price   int32             `json:"price,omitempty"`
	}

	FeatureInputArgs struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		DisplayName string `json:"display_name,omitempty"`
	}

	TrialInputArgs struct {
		TimeUnitName string  `json:"time_unit_name,omitempty"`
		TimeLength   float64 `json:"time_length,omitempty"`
		Credit       float64 `json:"credit,omitempty"`
	}

	DiscountInputArgs struct {
		DiscountType string  `json:"discount_type,omitempty"`
		Amount       float64 `json:"amount,omitempty"`
	}

	LimitInputArgs struct {
		Feature *FeatureInputArgs `json:"feature,omitempty"`
		Limit   float64           `json:"limit,omitempty"`
	}

	CreatePricePlanArgs struct {
		CouponName        string                      `json:"coupon_name,omitempty"`
		VendorId          int32                       `json:"vendor_id,omitempty"`
		Description       string                      `json:"description,omitempty"`
		Period            string                      `json:"period,omitempty"`
		BasePrice         int32                       `json:"base_price,omitempty"`
		AddOns            []AddOnInputArgs            `json:"add_ons,omitempty"`
		Tags              []string                    `json:"tags,omitempty"`
		Features          []FeatureInputArgs          `json:"features,omitempty"`
		Trial             *TrialInputArgs             `json:"trial,omitempty"`
		Discount          *DiscountInputArgs          `json:"discount,omitempty"`
		DisplayName       string                      `json:"display_name,omitempty"`
		Name              string                      `json:"name,omitempty"`
		Limits            []LimitInputArgs            `json:"limits,omitempty"`
		MeteredComponents []MeteredComponentInputArgs `json:"metered_components,omitempty"`
	}

	MeteredComponentInputArgs struct {
		MeterId     int32                 `json:"meter_id,omitempty"`
		PriceScheme *PriceSchemeInputArgs `json:"price_scheme,omitempty"`
		MeterName   string                `json:"meter_name,omitempty"`
	}

	PriceSchemeInputArgs struct {
		TimeUnitName string           `json:"time_unit_name,omitempty"`
		UnitName     string           `json:"unit_name,omitempty"`
		SchemeType   string           `json:"scheme_type,omitempty"`
		Prices       []PriceInputArgs `json:"prices,omitempty"`
		PriceList    []interface{}    `json:"price_list,omitempty"`
	}

	PriceInputArgs struct {
		Cap   float64 `json:"cap,omitempty"`
		Price float64 `json:"price,omitempty"`
	}

	UpdatePricePlanArgs struct {
		CouponName        string                      `json:"coupon_name,omitempty"`
		VendorId          int32                       `json:"vendor_id,omitempty"`
		Description       string                      `json:"description,omitempty"`
		Period            string                      `json:"period,omitempty"`
		BasePrice         int32                       `json:"base_price,omitempty"`
		AddOns            []AddOnInputArgs            `json:"add_ons,omitempty"`
		Tags              []string                    `json:"tags,omitempty"`
		Features          []FeatureInputArgs          `json:"features,omitempty"`
		Trial             *TrialInputArgs             `json:"trial,omitempty"`
		Discount          *DiscountInputArgs          `json:"discount,omitempty"`
		DisplayName       string                      `json:"display_name,omitempty"`
		Name              string                      `json:"name,omitempty"`
		Limits            []LimitInputArgs            `json:"limits,omitempty"`
		MeteredComponents []MeteredComponentInputArgs `json:"metered_components,omitempty"`
	}

	PricePlan swagger.PricePlan

	pricePlansAPI struct {
		impl *swagger.APIClient
		ctx  func() context.Context
	}
)

// Create creates Price Plan.
func (api *pricePlansAPI) Create(body CreatePricePlanArgs) (PricePlan, *http.Response, error) {
	implAddOns, implFeatures, implTrial, implDiscount, implLimits, implComponents :=
		pricePlanConvert(body.AddOns, body.Features, body.Trial, body.Discount, body.Limits, body.MeteredComponents)
	implCreatePricePlanArgs := swagger.CreatePricePlanArgs{
		CouponName:        body.CouponName,
		VendorId:          body.VendorId,
		Description:       body.Description,
		Period:            body.Period,
		BasePrice:         body.BasePrice,
		AddOns:            implAddOns,
		Tags:              body.Tags,
		Features:          implFeatures,
		Trial:             implTrial,
		Discount:          implDiscount,
		DisplayName:       body.DisplayName,
		Name:              body.Name,
		Limits:            implLimits,
		MeteredComponents: implComponents,
	}
	implPricePlan, resp, err := api.impl.PricePlansApi.PricePlansPost(
		api.ctx(), implCreatePricePlanArgs)
	pricePlan := implPricePlanToPricePlan(implPricePlan)
	return pricePlan, resp, err
}

// Retrieve fetches an existing price plan.
func (api *pricePlansAPI) Retrieve(pricePlanName string) (PricePlan, *http.Response, error) {
	implPricePlan, resp, err := api.impl.PricePlansApi.PricePlansPricePlanNameGet(
		api.ctx(), pricePlanName)
	pricePlan := implPricePlanToPricePlan(implPricePlan)
	return pricePlan, resp, err
}

// Update updates a price plan by its unique name.
func (api *pricePlansAPI) Update(pricePlanName string, body UpdatePricePlanArgs) (PricePlan, *http.Response, error) {
	implAddOns, implFeatures, implTrial, implDiscount, implLimits, implComponents :=
		pricePlanConvert(body.AddOns, body.Features, body.Trial, body.Discount, body.Limits, body.MeteredComponents)
	implUpdatePricePlanArgs := swagger.UpdatePricePlanArgs{
		CouponName:        body.CouponName,
		VendorId:          body.VendorId,
		Description:       body.Description,
		Period:            body.Period,
		BasePrice:         body.BasePrice,
		AddOns:            implAddOns,
		Tags:              body.Tags,
		Features:          implFeatures,
		Trial:             implTrial,
		Discount:          implDiscount,
		DisplayName:       body.DisplayName,
		Name:              body.Name,
		Limits:            implLimits,
		MeteredComponents: implComponents,
	}
	implPricePlan, resp, err := api.impl.PricePlansApi.PricePlansPricePlanNamePut(
		api.ctx(), implUpdatePricePlanArgs, pricePlanName)
	pricePlan := implPricePlanToPricePlan(implPricePlan)
	return pricePlan, resp, err
}

// List retrieves all price plans for a given vendor.
func (api *pricePlansAPI) List() ([]PricePlan, *http.Response, error) {
	implPricePlans, resp, err := api.impl.PricePlansApi.PricePlansGet(api.ctx())
	var pricePlans []PricePlan
	for _, implPricePlan := range implPricePlans {
		pricePlans = append(pricePlans, implPricePlanToPricePlan(implPricePlan))
	}
	return pricePlans, resp, err
}

// Delete deletes a price plan by its unique name.
func (api *pricePlansAPI) Delete(pricePlanName string) (*http.Response, error) {
	return api.impl.PricePlansApi.PricePlansPricePlanNameDelete(
		api.ctx(), pricePlanName)
}

func newPricePlansAPI(impl *swagger.APIClient, ctx func() context.Context) *pricePlansAPI {
	return &pricePlansAPI{impl: impl, ctx: ctx}
}

// Convert a Swagger PricePlan struct to our PricePlan struct
func implPricePlanToPricePlan(implPricePlan swagger.PricePlan) PricePlan {
	var pricePlan PricePlan
	if &implPricePlan != nil {
		pricePlan = PricePlan{
			Name:              implPricePlan.Name,
			DisplayName:       implPricePlan.DisplayName,
			Description:       implPricePlan.Description,
			BasePrice:         implPricePlan.BasePrice,
			Period:            implPricePlan.Period,
			Coupon:            implPricePlan.Coupon,
			MeteredComponents: implPricePlan.MeteredComponents,
			Discount:          implPricePlan.Discount,
			Features:          implPricePlan.Features,
			Tags:              implPricePlan.Tags,
			Trial:             implPricePlan.Trial,
		}
	}
	return pricePlan
}

// Convert various price plan fields for create/update... holy heck
func pricePlanConvert(addons []AddOnInputArgs, features []FeatureInputArgs, trial *TrialInputArgs, discount *DiscountInputArgs, limits []LimitInputArgs, components []MeteredComponentInputArgs) (
	implAddOns []swagger.AddOnInputArgs, implFeatures []swagger.FeatureInputArgs, implTrial *swagger.TrialInputArgs, implDiscount *swagger.DiscountInputArgs, implLimits []swagger.LimitInputArgs,
	implComponents []swagger.MeteredComponentInputArgs) {
	if addons != nil {
		for _, addon := range addons {
			var implFeature swagger.FeatureInputArgs
			if addon.Feature != nil {
				implFeature = swagger.FeatureInputArgs{
					Name:        addon.Feature.Name,
					Description: addon.Feature.Description,
					DisplayName: addon.Feature.DisplayName,
				}
			}
			implAddOns = append(implAddOns, swagger.AddOnInputArgs{
				Feature: &implFeature,
				Price:   addon.Price,
			})
		}
	}
	if features != nil {
		for _, feature := range features {
			implFeatures = append(implFeatures, swagger.FeatureInputArgs{
				Name:        feature.Name,
				Description: feature.Description,
				DisplayName: feature.DisplayName,
			})
		}
	}
	if trial != nil {
		implTrial = &swagger.TrialInputArgs{
			TimeUnitName: trial.TimeUnitName,
			TimeLength:   trial.TimeLength,
			Credit:       trial.Credit,
		}
	}
	if discount != nil {
		implDiscount = &swagger.DiscountInputArgs{
			DiscountType: discount.DiscountType,
			Amount:       discount.Amount,
		}
	}
	if limits != nil {
		for _, limit := range limits {
			var implFeature swagger.FeatureInputArgs
			if limit.Feature != nil {
				implFeature = swagger.FeatureInputArgs{
					Name:        limit.Feature.Name,
					Description: limit.Feature.Description,
					DisplayName: limit.Feature.DisplayName,
				}
			}
			implLimits = append(implLimits, swagger.LimitInputArgs{
				Feature: &implFeature,
				Limit:   limit.Limit,
			})
		}
	}
	if components != nil {
		for _, component := range components {
			var implPriceSchemeInputArgs swagger.PriceSchemeInputArgs
			if component.PriceScheme != nil {
				var implPriceInputArgs []swagger.PriceInputArgs
				for _, price := range component.PriceScheme.Prices {
					implPriceInputArgs = append(implPriceInputArgs, swagger.PriceInputArgs{
						Cap:   price.Cap,
						Price: price.Price,
					})
				}
				implPriceSchemeInputArgs = swagger.PriceSchemeInputArgs{
					TimeUnitName: component.PriceScheme.TimeUnitName,
					UnitName:     component.PriceScheme.UnitName,
					SchemeType:   component.PriceScheme.SchemeType,
					Prices:       implPriceInputArgs,
				}

				if component.PriceScheme.PriceList != nil {
					implPriceSchemeInputArgs = swagger.PriceSchemeInputArgs{
						TimeUnitName: component.PriceScheme.TimeUnitName,
						UnitName:     component.PriceScheme.UnitName,
						SchemeType:   component.PriceScheme.SchemeType,
						PriceList:    component.PriceScheme.PriceList,
					}
				}

			}
			implComponents = append(implComponents, swagger.MeteredComponentInputArgs{
				MeterId:     component.MeterId,
				PriceScheme: &implPriceSchemeInputArgs,
				MeterName:   component.MeterName,
			})
		}
	}
	return
}
