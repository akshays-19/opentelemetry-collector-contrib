// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oracledbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"net"
	"net/url"
	"strconv"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	go_ora "github.com/sijms/go-ora/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver/internal/metadata"
)

// NewFactory creates a new Oracle receiver factory.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createReceiverFunc(func(dataSourceName string) (*sql.DB, error) {
			return sql.Open("oracle", dataSourceName)
		}, newDbClient), metadata.MetricsStability),
		receiver.WithLogs(createLogsReceiverFunc(func(dataSourceName string) (*sql.DB, error) {
			return sql.Open("oracle", dataSourceName)
		}, newDbClient), metadata.MetricsStability))
}

func createDefaultConfig() component.Config {
	cfg := scraperhelper.NewDefaultControllerConfig()
	cfg.CollectionInterval = 10 * time.Second

	return &Config{
		ControllerConfig:     cfg,
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
		TopQueryCollection: TopQueryCollection{
			Enabled:             false,
			MaxQuerySampleCount: 1000,
			TopQueryCount:       200,
			QueryCacheSize:      5000,
		},
		QuerySample: QuerySample{
			Enabled: false,
		},
	}
}

type sqlOpenerFunc func(dataSourceName string) (*sql.DB, error)

func createReceiverFunc(sqlOpenerFunc sqlOpenerFunc, clientProviderFunc clientProviderFunc) receiver.CreateMetricsFunc {
	return func(
		_ context.Context,
		settings receiver.Settings,
		cfg component.Config,
		consumer consumer.Metrics,
	) (receiver.Metrics, error) {
		sqlCfg := cfg.(*Config)
		metricsBuilder := metadata.NewMetricsBuilder(sqlCfg.MetricsBuilderConfig, settings)

		instanceName, err := getInstanceName(getDataSource(*sqlCfg))
		if err != nil {
			return nil, err
		}

		mp, err := newScraper(metricsBuilder, sqlCfg.MetricsBuilderConfig, sqlCfg.ControllerConfig, settings.Logger, func() (*sql.DB, error) {
			return sqlOpenerFunc(getDataSource(*sqlCfg))
		}, clientProviderFunc, instanceName)
		if err != nil {
			return nil, err
		}
		opt := scraperhelper.AddScraper(metadata.Type, mp)

		return scraperhelper.NewMetricsController(
			&sqlCfg.ControllerConfig,
			settings,
			consumer,
			opt,
		)
	}
}

func createLogsReceiverFunc(sqlOpenerFunc sqlOpenerFunc, clientProviderFunc clientProviderFunc) receiver.CreateLogsFunc {
	return func(
		_ context.Context,
		settings receiver.Settings,
		cfg component.Config,
		logsConsumer consumer.Logs,
	) (receiver.Logs, error) {
		sqlCfg := cfg.(*Config)

		if !sqlCfg.TopQueryCollection.Enabled && !sqlCfg.QuerySample.Enabled {
			settings.TelemetrySettings.Logger.Debug("TopQueryCollection and QuerySample are not enabled for Oracle receiver.Skipping Log scrapper")
			return nil, nil
		}

		instanceName, err := getInstanceName(getDataSource(*sqlCfg))
		if err != nil {
			return nil, err
		}

		hostName, hostNameErr := getHostName(getDataSource(*sqlCfg))
		if hostNameErr != nil {
			return nil, hostNameErr
		}

		cacheSize := sqlCfg.QueryCacheSize
		metricCache, err := lru.New[string, map[string]int64](cacheSize)
		if err != nil {
			settings.TelemetrySettings.Logger.Error("Failed to create LRU cache, skipping the current scraper", zap.Error(err))
			return nil, err
		}

		mp, err := newLogsScraper(sqlCfg.MetricsBuilderConfig, sqlCfg.ControllerConfig, settings.TelemetrySettings.Logger, func() (*sql.DB, error) {
			return sqlOpenerFunc(getDataSource(*sqlCfg))
		}, clientProviderFunc, instanceName, metricCache, sqlCfg.TopQueryCollection, sqlCfg.QuerySample, hostName)
		if err != nil {
			return nil, err
		}
		// adding a logs scraper is still not properly implemented in the helper, so we need to c&p some of that code here
		// to make a logs scraper work
		f := scraper.NewFactory(metadata.Type, nil,
			scraper.WithLogs(func(context.Context, scraper.Settings, component.Config) (scraper.Logs, error) {
				return mp, nil
			}, component.StabilityLevelAlpha))
		opt := scraperhelper.AddFactoryWithConfig(f, nil)

		return scraperhelper.NewLogsController(
			&sqlCfg.ControllerConfig,
			settings,
			logsConsumer,
			opt,
		)
	}
}

func getDataSource(cfg Config) string {
	if cfg.DataSource != "" {
		return cfg.DataSource
	}

	// Don't need to worry about errors here as config validation already checked.
	host, portStr, _ := net.SplitHostPort(cfg.Endpoint)
	port, _ := strconv.ParseInt(portStr, 10, 32)

	return go_ora.BuildUrl(host, int(port), cfg.Service, cfg.Username, cfg.Password, nil)
}

func getInstanceName(datasource string) (string, error) {
	datasourceURL, err := url.Parse(datasource)
	if err != nil {
		return "", err
	}

	instanceName := datasourceURL.Host + datasourceURL.Path
	return instanceName, nil
}

func getHostName(datasource string) (string, error) {
	datasourceURL, err := url.Parse(datasource)
	if err != nil {
		return "", err
	}
	return datasourceURL.Host, nil
}
