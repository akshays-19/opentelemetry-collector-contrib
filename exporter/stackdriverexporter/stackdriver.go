// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package stackdriverexporter contains the wrapper for OpenTelemetry-Stackdriver
// exporter to be used in OpenTelemetry-Service.
package stackdriverexporter

import (
	"context"
	"fmt"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"

	"github.com/open-telemetry/opentelemetry-service/consumer"
	"github.com/open-telemetry/opentelemetry-service/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-service/exporter/exporterwrapper"
)

type stackdriverConfig struct {
	ProjectID     string `mapstructure:"project,omitempty"`
	EnableTracing bool   `mapstructure:"enable-tracing,omitempty"`
	EnableMetrics bool   `mapstructure:"enable-metrics,omitempty"`
	MetricPrefix  string `mapstructure:"metric-prefix,omitempty"`
}

// TODO: Add metrics support to the exporterwrapper.
type stackdriverExporter struct {
	exporter *stackdriver.Exporter
}

var _ consumer.MetricsConsumer = (*stackdriverExporter)(nil)

func newStackdriverTraceExporter(ProjectID, MetricPrefix string) (consumer.TraceConsumer, func() error, error) {
	sde, serr := newStackdriverExporter(ProjectID, MetricPrefix)
	if serr != nil {
		return nil, nil, fmt.Errorf("cannot configure Stackdriver Trace exporter: %v", serr)
	}

	tExp, err := exporterwrapper.NewExporterWrapper("stackdriver_trace", "ocservice.exporter.Stackdriver.ConsumeTraceData", sde)
	if err != nil {
		return nil, nil, err
	}
	// TODO: Examine "contrib.go.opencensus.io/exporter/stackdriver" to see
	// if trace.ExportSpan was constraining and if perhaps the Stackdriver
	// upload can use the context and information from the Node.

	doneFn := func() error {
		sde.Flush()
		return nil
	}

	return tExp, doneFn, nil
}

func newStackdriverMetricsExporter(ProjectID, MetricPrefix string) (consumer.MetricsConsumer, func() error, error) {
	sde, serr := newStackdriverExporter(ProjectID, MetricPrefix)
	if serr != nil {
		return nil, nil, fmt.Errorf("cannot configure Stackdriver metric exporter: %v", serr)
	}

	mExp := &stackdriverExporter{
		exporter: sde,
	}

	doneFn := func() error {
		sde.Flush()
		return nil
	}

	return mExp, doneFn, nil
}

func newStackdriverExporter(ProjectID, MetricPrefix string) (*stackdriver.Exporter, error) {
	// TODO:  For each ProjectID, create a different exporter
	// or at least a unique Stackdriver client per ProjectID.

	return stackdriver.NewExporter(stackdriver.Options{
		// If the project ID is an empty string, it will be set by default based on
		// the project this is running on in GCP.
		ProjectID: ProjectID,

		MetricPrefix: MetricPrefix,

		// Stackdriver Metrics mandates a minimum of 60 seconds for
		// reporting metrics. We have to enforce this as per the advisory
		// at https://cloud.google.com/monitoring/custom-metrics/creating-metrics#writing-ts
		// which says:
		//
		// "If you want to write more than one point to the same time series, then use a separate call
		//  to the timeSeries.create method for each point. Don't make the calls faster than one time per
		//  minute. If you are adding data points to different time series, then there is no rate limitation."
		BundleDelayThreshold: 61 * time.Second,
	})
}

func (sde *stackdriverExporter) ConsumeMetricsData(ctx context.Context, md consumerdata.MetricsData) error {
	ctx, span := trace.StartSpan(ctx,
		"opentelemetry.service.exporter.stackdriver.ExportMetricsData",
		trace.WithSampler(trace.NeverSample()))
	defer span.End()

	err := sde.exporter.ExportMetricsProto(ctx, md.Node, md.Resource, md.Metrics)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
	}

	return err
}
