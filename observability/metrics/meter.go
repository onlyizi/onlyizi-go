package metrics

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"

	metricapi "go.opentelemetry.io/otel/metric"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

var meterProvider *metricsdk.MeterProvider

func Init() error {

	exporter, err := prometheus.New()
	if err != nil {
		return err
	}

	meterProvider = metricsdk.NewMeterProvider(
		metricsdk.WithReader(exporter),
	)

	otel.SetMeterProvider(meterProvider)

	return nil
}

func Meter(name string) metricapi.Meter {
	return otel.Meter(name)
}
