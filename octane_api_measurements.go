package octane

import (
	"context"
	"net/http"
	"time"

	"github.com/getoctane/octane-go/internal/swagger"
)

type (
	Measurement swagger.Measurement

	measurementsAPI struct {
		impl *swagger.APIClient
		ctx  func() context.Context
	}
)

func newMeasurementsAPI(impl *swagger.APIClient, ctx func() context.Context) *measurementsAPI {
	return &measurementsAPI{impl: impl, ctx: ctx}
}

// Create sends a measurement.
func (api *measurementsAPI) Create(body Measurement) (Measurement, *http.Response, error) {
	if body.Time.IsZero() {
		body.Time = time.Now()
	}
	implMeasurement := swagger.Measurement{
		Value:        body.Value,
		Labels:       body.Labels,
		Time:         body.Time,
		MeterName:    body.MeterName,
		ResetTotal:   body.ResetTotal,
		CustomerName: body.CustomerName,
	}
	implMeasurementReturn, resp, err := api.impl.MeasurementsApi.MeasurementsPost(
		api.ctx(), implMeasurement)
	measurement := implMeasurementToMeasurement(&implMeasurementReturn)
	return measurement, resp, err
}

// CreateMulti sends multiple measurements.
func (api *measurementsAPI) CreateMulti(body []Measurement) ([]Measurement, *http.Response, error) {
	var implMeasurements []swagger.Measurement
	for _, measurement := range body {
		if measurement.Time.IsZero() {
			measurement.Time = time.Now()
		}
		implMeasurements = append(implMeasurements, swagger.Measurement{
			Value:      measurement.Value,
			Labels:     measurement.Labels,
			Time:       measurement.Time,
			MeterName:  measurement.MeterName,
			ResetTotal: measurement.ResetTotal,
		})
	}
	implMeasurementsReturn, resp, err := api.impl.MeasurementsApi.MeasurementsMultiPost(
		api.ctx(), implMeasurements)
	var measurements []Measurement
	for _, implMeasurement := range implMeasurementsReturn {
		measurements = append(measurements, implMeasurementToMeasurement(&implMeasurement))
	}
	return measurements, resp, err
}

// Convert a Swagger Measurement struct to our Measurement struct
func implMeasurementToMeasurement(implMeasurement *swagger.Measurement) Measurement {
	var measurement Measurement
	if implMeasurement != nil {
		measurement = Measurement{
			Value:        implMeasurement.Value,
			Labels:       implMeasurement.Labels,
			Time:         implMeasurement.Time,
			MeterName:    implMeasurement.MeterName,
			ResetTotal:   implMeasurement.ResetTotal,
			CustomerName: implMeasurement.CustomerName,
		}
	}
	return measurement
}
