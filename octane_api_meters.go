package octane

import (
	"context"
	"net/http"

	"github.com/getoctane/octane-go/internal/swagger"
)

type (
	MeterInputArgs  swagger.MeterInputArgs
	UpdateMeterArgs swagger.UpdateMeterArgs
	Meter           swagger.Meter

	metersAPI struct {
		impl *swagger.APIClient
		ctx  func() context.Context
	}
)

func newMetersAPI(impl *swagger.APIClient, ctx func() context.Context) *metersAPI {
	return &metersAPI{impl: impl, ctx: ctx}
}

// Create creates a new Meter.
func (api *metersAPI) Create(body MeterInputArgs) (Meter, *http.Response, error) {
	implMeterInputArgs := swagger.MeterInputArgs{
		VendorId:       body.VendorId,
		Description:    body.Description,
		PrimaryLabels:  body.PrimaryLabels,
		MeterType:      body.MeterType,
		DisplayName:    body.DisplayName,
		IsIncremental:  body.IsIncremental,
		Name:           body.Name,
		UnitName:       body.UnitName,
		ExpectedLabels: body.ExpectedLabels,
	}
	implMeter, resp, err := api.impl.MetersApi.MetersPost(
		api.ctx(), implMeterInputArgs)
	meter := implMeterToMeter(implMeter)
	return meter, resp, err
}

// Retrieve fetches a meter by its unique name.
func (api *metersAPI) Retrieve(meterName string) (Meter, *http.Response, error) {
	implMeter, resp, err := api.impl.MetersApi.MetersMeterNameGet(
		api.ctx(), meterName)
	meter := implMeterToMeter(implMeter)
	return meter, resp, err
}

// Update updates a meter by its unique name.
func (api *metersAPI) Update(meterName string, body UpdateMeterArgs) (Meter, *http.Response, error) {
	implUpdateMeterArgs := swagger.UpdateMeterArgs{
		Description:    body.Description,
		DisplayName:    body.DisplayName,
	}
	implMeter, resp, err := api.impl.MetersApi.MetersMeterNamePut(
		api.ctx(), implUpdateMeterArgs, meterName)
	meter := implMeterToMeter(implMeter)
	return meter, resp, err
}

// List retrieves all meters for a given vendor.
func (api *metersAPI) List() ([]Meter, *http.Response, error) {
	implMeters, resp, err := api.impl.MetersApi.MetersGet(api.ctx())
	var meters []Meter
	for _, implMeter := range implMeters {
		meters = append(meters, implMeterToMeter(implMeter))
	}
	return meters, resp, err
}

// Delete deletes a meter by its unique name.
func (api *metersAPI) Delete(meterName string) (*http.Response, error) {
	return api.impl.MetersApi.MetersMeterNameDelete(
		api.ctx(), meterName)
}

// Convert a Swagger Meter struct to our Meter struct
func implMeterToMeter(implMeter swagger.Meter) Meter {
	var meter Meter
	if &implMeter != nil {
		meter = Meter{
			Name:           implMeter.Name,
			DisplayName:    implMeter.DisplayName,
			Description:    implMeter.Description,
			IsIncremental:  implMeter.IsIncremental,
			MeterType:      implMeter.MeterType,
			UnitName:       implMeter.UnitName,
			ExpectedLabels: implMeter.ExpectedLabels,
			PrimaryLabels:  implMeter.PrimaryLabels,
		}
	}
	return meter
}
