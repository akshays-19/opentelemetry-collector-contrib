// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oracledbreceiver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.opentelemetry.io/collector/pdata/plog"
	"strings"
	"testing"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver/internal/metadata"
)

func TestScraper_ErrorOnStart(t *testing.T) {
	scrpr := oracleScraper{
		dbProviderFunc: func() (*sql.DB, error) {
			return nil, errors.New("oops")
		},
	}
	err := scrpr.start(context.Background(), componenttest.NewNopHost())
	require.Error(t, err)
}

var queryResponses = map[string][]metricRow{
	statsSQL:        {{"NAME": enqueueDeadlocks, "VALUE": "18"}, {"NAME": exchangeDeadlocks, "VALUE": "88898"}, {"NAME": executeCount, "VALUE": "178878"}, {"NAME": parseCountTotal, "VALUE": "1999"}, {"NAME": parseCountHard, "VALUE": "1"}, {"NAME": userCommits, "VALUE": "187778888"}, {"NAME": userRollbacks, "VALUE": "1898979879789"}, {"NAME": physicalReads, "VALUE": "1887777"}, {"NAME": physicalReadsDirect, "VALUE": "31337"}, {"NAME": sessionLogicalReads, "VALUE": "189"}, {"NAME": cpuTime, "VALUE": "1887"}, {"NAME": pgaMemory, "VALUE": "1999887"}, {"NAME": dbBlockGets, "VALUE": "42"}, {"NAME": consistentGets, "VALUE": "78944"}},
	sessionCountSQL: {{"VALUE": "1"}},
	systemResourceLimitsSQL: {
		{"RESOURCE_NAME": "processes", "CURRENT_UTILIZATION": "3", "MAX_UTILIZATION": "10", "INITIAL_ALLOCATION": "100", "LIMIT_VALUE": "100"},
		{"RESOURCE_NAME": "locks", "CURRENT_UTILIZATION": "3", "MAX_UTILIZATION": "10", "INITIAL_ALLOCATION": "-1", "LIMIT_VALUE": "-1"},
	},
	tablespaceUsageSQL:     {{"TABLESPACE_NAME": "SYS", "USED_SPACE": "111288", "TABLESPACE_SIZE": "3518587", "BLOCK_SIZE": "8192"}},
	oracleQueryMetricsData: {{"APPLICATION_WAIT_TIME": "0", "BUFFER_GETS": "3997614", "CHILD_ADDRESS": "0000000074C6E830", "CHILD_NUMBER": "0", "CLUSTER_WAIT_TIME": "2000", "CONCURRENCY_WAIT_TIME": "30", "CPU_TIME": "39736887", "DIRECT_READS": "5", "DIRECT_WRITES": "10", "DISK_READS": "15", "ELAPSED_TIME": "99172810", "EXECUTIONS": "300413", "PHYSICAL_READ_BYTES": "400", "PHYSICAL_READ_REQUESTS": "500", "PHYSICAL_WRITE_BYTES": "18", "PHYSICAL_WRITE_REQUESTS": "150", "ROWS_PROCESSED": "399856", "SQL_FULLTEXT": "SELECT COUNT(*) AS TotalHighRateEmployees FROM ADMIN.\"Employee\" e \nJOIN ADMIN.\"Person\" p ON e.\"BusinessEntityID\" = \"BusId1234\" \nJOIN ADMIN.\"EmployeePayHistory\" eph ON e.\"BusinessEntityID\" = eph.\"BusinessEntityID\" WHERE eph.\"Rate\" \u003e \n(SELECT AVG(\"Rate\") FROM ADMIN.\"EmployeePayHistory\")", "SQL_ID": "fxk8aq3nds8aw", "USER_IO_WAIT_TIME": "220"}},
	oracleQueryPlanData:    {{"ACCESS_PREDICATES": "", "BYTES": "", "CARDINALITY": "", "CHILD_ADDRESS": "0000000074C6E830", "COST": "1", "CPU_COST": "", "DEPTH": "0", "ID": "0", "IO_COST": "", "OPERATION": "SELECT STATEMENT", "PARENT_ID": "", "POSITION": "1", "PROJECTION": ""}},
}

var cacheValue = map[string]int64{"APPLICATION_WAIT_TIME": 0, "BUFFER_GETS": 3808197, "CLUSTER_WAIT_TIME": 1000, "CONCURRENCY_WAIT_TIME": 20, "CPU_TIME": 29821063, "DIRECT_READS": 3, "DIRECT_WRITES": 6, "DISK_READS": 12, "ELAPSED_TIME": 38172810, "EXECUTIONS": 200413, "PHYSICAL_READ_BYTES": 300, "PHYSICAL_READ_REQUESTS": 100, "PHYSICAL_WRITE_BYTES": 12, "PHYSICAL_WRITE_REQUESTS": 120, "ROWS_PROCESSED": 200413, "USER_IO_WAIT_TIME": 200}
var planJsonText = "[{\"ACCESS_PREDICATES\":\"\",\"BYTES\":\"\",\"CARDINALITY\":\"\",\"COST\":\"1\",\"CPU_COST\":\"\",\"DEPTH\":\"0\",\"ID\":\"0\",\"IO_COST\":\"\",\"OPERATION\":\"SELECT STATEMENT\",\"PARENT_ID\":\"\",\"POSITION\":\"1\",\"PROJECTION\":\"\"}]"

func TestScraper_Scrape(t *testing.T) {
	tests := []struct {
		name       string
		dbclientFn func(db *sql.DB, s string, logger *zap.Logger) dbClient
		errWanted  string
	}{
		{
			name: "valid",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				return &fakeDbClient{
					Responses: [][]metricRow{
						queryResponses[s],
					},
				}
			},
		},
		{
			name: "bad tablespace usage",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				if s == tablespaceUsageSQL {
					return &fakeDbClient{Responses: [][]metricRow{
						{
							{},
						},
					}}
				}
				return &fakeDbClient{Responses: [][]metricRow{
					queryResponses[s],
				}}
			},
			errWanted: `failed to parse int64 for OracledbTablespaceSizeUsage, value was : strconv.ParseInt: parsing "": invalid syntax`,
		},
		{
			name: "no limit on tablespace",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				if s == tablespaceUsageSQL {
					return &fakeDbClient{Responses: [][]metricRow{
						{
							{"TABLESPACE_NAME": "SYS", "TABLESPACE_SIZE": "1024", "USED_SPACE": "111288", "BLOCK_SIZE": "8192"},
							{"TABLESPACE_NAME": "FOO", "TABLESPACE_SIZE": "", "USED_SPACE": "111288", "BLOCK_SIZE": "8192"},
						},
					}}
				}
				return &fakeDbClient{Responses: [][]metricRow{
					queryResponses[s],
				}}
			},
		},
		{
			name: "bad value on tablespace",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				if s == tablespaceUsageSQL {
					return &fakeDbClient{Responses: [][]metricRow{
						{
							{"TABLESPACE_NAME": "SYS", "TABLESPACE_SIZE": "1024", "USED_SPACE": "111288", "BLOCK_SIZE": "8192"},
							{"TABLESPACE_NAME": "FOO", "TABLESPACE_SIZE": "ert", "USED_SPACE": "111288", "BLOCK_SIZE": "8192"},
						},
					}}
				}
				return &fakeDbClient{Responses: [][]metricRow{
					queryResponses[s],
				}}
			},
			errWanted: `failed to parse int64 for OracledbTablespaceSizeLimit, value was ert: strconv.ParseInt: parsing "ert": invalid syntax`,
		},
		{
			name: "Empty block size",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				if s == tablespaceUsageSQL {
					return &fakeDbClient{Responses: [][]metricRow{
						{
							{"TABLESPACE_NAME": "SYS", "TABLESPACE_SIZE": "1024", "USED_SPACE": "111288", "BLOCK_SIZE": ""},
						},
					}}
				}
				return &fakeDbClient{Responses: [][]metricRow{
					queryResponses[s],
				}}
			},
			errWanted: `failed to parse int64 for OracledbBlockSize, value was : strconv.ParseInt: parsing "": invalid syntax`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := metadata.DefaultMetricsBuilderConfig()
			cfg.Metrics.OracledbConsistentGets.Enabled = true
			cfg.Metrics.OracledbDbBlockGets.Enabled = true

			scrpr := oracleScraper{
				logger: zap.NewNop(),
				mb:     metadata.NewMetricsBuilder(cfg, receivertest.NewNopSettings(metadata.Type)),
				dbProviderFunc: func() (*sql.DB, error) {
					return nil, nil
				},
				clientProviderFunc:   test.dbclientFn,
				id:                   component.ID{},
				metricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
			}
			err := scrpr.start(context.Background(), componenttest.NewNopHost())
			defer func() {
				assert.NoError(t, scrpr.shutdown(context.Background()))
			}()
			require.NoError(t, err)
			m, err := scrpr.scrape(context.Background())
			if test.errWanted != "" {
				require.True(t, scrapererror.IsPartialScrapeError(err))
				require.EqualError(t, err, test.errWanted)
			} else {
				require.NoError(t, err)
				assert.Equal(t, 18, m.ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics().Len())
			}
			name, ok := m.ResourceMetrics().At(0).Resource().Attributes().Get("oracledb.instance.name")
			assert.True(t, ok)
			assert.Empty(t, name.Str())
			var found pmetric.Metric
			for i := 0; i < m.ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics().Len(); i++ {
				metric := m.ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics().At(i)
				if metric.Name() == "oracledb.consistent_gets" {
					found = metric
					break
				}
			}
			assert.Equal(t, int64(78944), found.Sum().DataPoints().At(0).IntValue())
		})
	}
}

var samplesQueryResponses = map[string][]metricRow{
	samplesQuery: {{"S.MACHINE": "TEST-MACHINE", "S.USERNAME": "ADMIN", "S.SCHEMANAME": "ADMIN", "S.SQL_ID": "48bc50b6fuz4y",
		"S.SQL_CHILD_NUMBER": "0", "S.SID": "675", "S.SERIAL#": "51295", "Q.SQL_FULLTEXT": "test_query", "S.OSUSER": "test-user", "S.PROCESS": "1115",
		"S.PORT": "54440", "S.PROGRAM": "Oracle SQL Developer for VS Code", "S.MODULE": "Oracle SQL Developer for VS Code", "S.STATUS": "ACTIVE", "S.STATE": "WAITED KNOWN TIME", "Q.PLAN_HASH_VALUE": "4199919568", "DURATION_SEC": "1"}},
}

func TestSamplesQuery(t *testing.T) {
	tests := []struct {
		name       string
		dbclientFn func(db *sql.DB, s string, logger *zap.Logger) dbClient
		errWanted  string
	}{
		{
			name: "valid",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				return &fakeDbClient{
					Responses: [][]metricRow{
						samplesQueryResponses[s],
					},
				}
			},
		},
		{
			name: "bad samples data",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				if s == samplesQuery {
					return &fakeDbClient{Responses: [][]metricRow{
						{
							{},
						},
					}}
				}
				return &fakeDbClient{Responses: [][]metricRow{
					samplesQueryResponses[s],
				}}
			},
			errWanted: `failed to parse int64 for Duration, value was : strconv.ParseInt: parsing "": invalid syntax`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			scrpr := oracleScraper{
				logger: zap.NewNop(),
				dbProviderFunc: func() (*sql.DB, error) {
					return nil, nil
				},
				clientProviderFunc: test.dbclientFn,
				id:                 component.ID{},
			}
			err := scrpr.start(context.Background(), componenttest.NewNopHost())
			defer func() {
				assert.NoError(t, scrpr.shutdown(context.Background()))
			}()
			require.NoError(t, err)
			m := plog.NewLogs()
			err = scrpr.collectQuerySamples(context.Background(), m)

			if test.errWanted != "" {
				//require.True(t, scrapererror.IsPartialScrapeError(err))
				require.EqualError(t, err, test.errWanted)
			} else {
				require.NoError(t, err)
				assert.Equal(t, 20, m.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Attributes().Len())
			}
			name, ok := m.ResourceLogs().At(0).Resource().Attributes().Get("oracledb.instance.name")
			assert.True(t, ok)
			assert.Empty(t, name.Str())
		})
	}
}

func TestScraper_ScrapeLogs(t *testing.T) {
	tests := []struct {
		name       string
		dbclientFn func(db *sql.DB, s string, logger *zap.Logger) dbClient
		errWanted  string
	}{
		{
			name: "valid collection",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				if strings.Contains(s, "V$SQL_PLAN") {
					return &fakeDbClient{
						Responses: [][]metricRow{
							queryResponses[oracleQueryPlanData],
						},
					}

				} else {
					return &fakeDbClient{
						Responses: [][]metricRow{
							queryResponses[oracleQueryMetricsData],
						},
					}
				}
			},
		}, {
			name: "No metrics collected",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				return &fakeDbClient{
					Responses: [][]metricRow{
						nil,
					},
				}
			},
			errWanted: `no data returned from oracleQueryMetricsClient`,
		}, {
			name: "Error on collecting metrics",
			dbclientFn: func(_ *sql.DB, s string, _ *zap.Logger) dbClient {
				return &fakeDbClient{
					Responses: [][]metricRow{
						nil,
					},
					Err: errors.New("Mock error"),
				}
			},
			errWanted: fmt.Sprintf("error executing %s: %s", oracleQueryMetricsData, "Mock error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := metadata.DefaultMetricsBuilderConfig()
			cfg.Metrics.OracledbConsistentGets.Enabled = true
			cfg.Metrics.OracledbDbBlockGets.Enabled = true
			lruCache, _ := lru.New[string, map[string]int64](500)
			lruCache.Add("fxk8aq3nds8aw:0", cacheValue)

			scrpr := oracleScraper{
				logger: zap.NewNop(),
				mb:     metadata.NewMetricsBuilder(cfg, receivertest.NewNopSettings(receivertest.NopType)),
				dbProviderFunc: func() (*sql.DB, error) {
					return nil, nil
				},
				clientProviderFunc:   test.dbclientFn,
				id:                   component.ID{},
				metricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
				metricCache:          lruCache,
				topQueryCount:        200,
			}

			scrpr.metricsBuilderConfig.Metrics.OracledbQueryElapsedTime.Enabled = true
			scrpr.metricsBuilderConfig.Metrics.OracledbQueryExecutions.Enabled = true
			scrpr.metricsBuilderConfig.Metrics.OracledbQueryDiskReads.Enabled = true
			scrpr.metricsBuilderConfig.Metrics.OracledbQueryDirectWrites.Enabled = true
			scrpr.metricsBuilderConfig.Metrics.OracledbQueryBufferGets.Enabled = true

			err := scrpr.start(context.Background(), componenttest.NewNopHost())
			defer func() {
				assert.NoError(t, scrpr.shutdown(context.Background()))
			}()
			require.NoError(t, err)

			logs := plog.NewLogs()
			err = scrpr.collectTopNMetricData(context.Background(), logs)

			if test.errWanted != "" {
				require.EqualError(t, err, test.errWanted)
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().Len(), "Query metrics has not been added to LogRecords")
				assert.Equal(t, 9, logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Attributes().Len(), "Metric data missing in LogRecord")
				elapsedTimeValue, _ := logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Attributes().Get("oracledb.query.elapsed_time")
				assert.Equal(t, 61, int(elapsedTimeValue.Double()), "Metric value calculation error")
				executionsValue, _ := logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Attributes().Get("oracledb.query.executions")
				assert.Equal(t, 100000, int(executionsValue.Int()), "Metric value calculation error")

				systemName, _ := logs.ResourceLogs().At(0).Resource().Attributes().Get("db.system.name")
				assert.Equal(t, systemName.AsString(), "oracle.db", "db.system.name not set correctly")

				queryText, _ := logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Attributes().Get("db.query.text")
				assert.False(t, strings.Contains(queryText.AsString(), "BusId1234"), "Obfuscation failure")

				planText, _ := logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Attributes().Get("oracledb.query_plan")
				assert.Equal(t, planJsonText, planText.AsString(), "Plan text not accurate")

			}

		})
	}

}

func TestTopNCollection_ShouldExitIfDisabled(t *testing.T) {
	cfg := metadata.DefaultMetricsBuilderConfig()
	lruCache, _ := lru.New[string, map[string]int64](500)
	scrpr := oracleScraper{
		logger: zap.NewNop(),
		mb:     metadata.NewMetricsBuilder(cfg, receivertest.NewNopSettings(receivertest.NopType)),
		dbProviderFunc: func() (*sql.DB, error) {
			return nil, nil
		},
		clientProviderFunc:   func(_ *sql.DB, s string, _ *zap.Logger) dbClient { return &fakeDbClient{} },
		id:                   component.ID{},
		metricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
		metricCache:          lruCache,
		topQueryCount:        200,
	}

	scrpr.metricsBuilderConfig.Metrics.OracledbQueryElapsedTime.Enabled = false

	scrpr.start(context.Background(), componenttest.NewNopHost())
	logs := plog.NewLogs()
	err2 := scrpr.collectTopNMetricData(context.Background(), logs)
	assert.Nil(t, err2, "Collection should gracefully exit if OracledbQueryElapsedTime is disabled in config")
	assert.Zero(t, logs.ResourceLogs().Len(), "Collection should not happen if OracledbQueryElapsedTime is disabled in config")
}
