// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for mongodbatlas metrics.
type MetricsConfig struct {
	MongodbatlasDbCounts                                  MetricConfig `mapstructure:"mongodbatlas.db.counts"`
	MongodbatlasDbSize                                    MetricConfig `mapstructure:"mongodbatlas.db.size"`
	MongodbatlasDiskPartitionIopsAverage                  MetricConfig `mapstructure:"mongodbatlas.disk.partition.iops.average"`
	MongodbatlasDiskPartitionIopsMax                      MetricConfig `mapstructure:"mongodbatlas.disk.partition.iops.max"`
	MongodbatlasDiskPartitionLatencyAverage               MetricConfig `mapstructure:"mongodbatlas.disk.partition.latency.average"`
	MongodbatlasDiskPartitionLatencyMax                   MetricConfig `mapstructure:"mongodbatlas.disk.partition.latency.max"`
	MongodbatlasDiskPartitionQueueDepth                   MetricConfig `mapstructure:"mongodbatlas.disk.partition.queue.depth"`
	MongodbatlasDiskPartitionSpaceAverage                 MetricConfig `mapstructure:"mongodbatlas.disk.partition.space.average"`
	MongodbatlasDiskPartitionSpaceMax                     MetricConfig `mapstructure:"mongodbatlas.disk.partition.space.max"`
	MongodbatlasDiskPartitionThroughput                   MetricConfig `mapstructure:"mongodbatlas.disk.partition.throughput"`
	MongodbatlasDiskPartitionUsageAverage                 MetricConfig `mapstructure:"mongodbatlas.disk.partition.usage.average"`
	MongodbatlasDiskPartitionUsageMax                     MetricConfig `mapstructure:"mongodbatlas.disk.partition.usage.max"`
	MongodbatlasDiskPartitionUtilizationAverage           MetricConfig `mapstructure:"mongodbatlas.disk.partition.utilization.average"`
	MongodbatlasDiskPartitionUtilizationMax               MetricConfig `mapstructure:"mongodbatlas.disk.partition.utilization.max"`
	MongodbatlasProcessAsserts                            MetricConfig `mapstructure:"mongodbatlas.process.asserts"`
	MongodbatlasProcessBackgroundFlush                    MetricConfig `mapstructure:"mongodbatlas.process.background_flush"`
	MongodbatlasProcessCacheIo                            MetricConfig `mapstructure:"mongodbatlas.process.cache.io"`
	MongodbatlasProcessCacheRatio                         MetricConfig `mapstructure:"mongodbatlas.process.cache.ratio"`
	MongodbatlasProcessCacheSize                          MetricConfig `mapstructure:"mongodbatlas.process.cache.size"`
	MongodbatlasProcessConnections                        MetricConfig `mapstructure:"mongodbatlas.process.connections"`
	MongodbatlasProcessCPUChildrenNormalizedUsageAverage  MetricConfig `mapstructure:"mongodbatlas.process.cpu.children.normalized.usage.average"`
	MongodbatlasProcessCPUChildrenNormalizedUsageMax      MetricConfig `mapstructure:"mongodbatlas.process.cpu.children.normalized.usage.max"`
	MongodbatlasProcessCPUChildrenUsageAverage            MetricConfig `mapstructure:"mongodbatlas.process.cpu.children.usage.average"`
	MongodbatlasProcessCPUChildrenUsageMax                MetricConfig `mapstructure:"mongodbatlas.process.cpu.children.usage.max"`
	MongodbatlasProcessCPUNormalizedUsageAverage          MetricConfig `mapstructure:"mongodbatlas.process.cpu.normalized.usage.average"`
	MongodbatlasProcessCPUNormalizedUsageMax              MetricConfig `mapstructure:"mongodbatlas.process.cpu.normalized.usage.max"`
	MongodbatlasProcessCPUUsageAverage                    MetricConfig `mapstructure:"mongodbatlas.process.cpu.usage.average"`
	MongodbatlasProcessCPUUsageMax                        MetricConfig `mapstructure:"mongodbatlas.process.cpu.usage.max"`
	MongodbatlasProcessCursors                            MetricConfig `mapstructure:"mongodbatlas.process.cursors"`
	MongodbatlasProcessDbDocumentRate                     MetricConfig `mapstructure:"mongodbatlas.process.db.document.rate"`
	MongodbatlasProcessDbOperationsRate                   MetricConfig `mapstructure:"mongodbatlas.process.db.operations.rate"`
	MongodbatlasProcessDbOperationsTime                   MetricConfig `mapstructure:"mongodbatlas.process.db.operations.time"`
	MongodbatlasProcessDbQueryExecutorScanned             MetricConfig `mapstructure:"mongodbatlas.process.db.query_executor.scanned"`
	MongodbatlasProcessDbQueryTargetingScannedPerReturned MetricConfig `mapstructure:"mongodbatlas.process.db.query_targeting.scanned_per_returned"`
	MongodbatlasProcessDbStorage                          MetricConfig `mapstructure:"mongodbatlas.process.db.storage"`
	MongodbatlasProcessGlobalLock                         MetricConfig `mapstructure:"mongodbatlas.process.global_lock"`
	MongodbatlasProcessIndexBtreeMissRatio                MetricConfig `mapstructure:"mongodbatlas.process.index.btree_miss_ratio"`
	MongodbatlasProcessIndexCounters                      MetricConfig `mapstructure:"mongodbatlas.process.index.counters"`
	MongodbatlasProcessJournalingCommits                  MetricConfig `mapstructure:"mongodbatlas.process.journaling.commits"`
	MongodbatlasProcessJournalingDataFiles                MetricConfig `mapstructure:"mongodbatlas.process.journaling.data_files"`
	MongodbatlasProcessJournalingWritten                  MetricConfig `mapstructure:"mongodbatlas.process.journaling.written"`
	MongodbatlasProcessMemoryUsage                        MetricConfig `mapstructure:"mongodbatlas.process.memory.usage"`
	MongodbatlasProcessNetworkIo                          MetricConfig `mapstructure:"mongodbatlas.process.network.io"`
	MongodbatlasProcessNetworkRequests                    MetricConfig `mapstructure:"mongodbatlas.process.network.requests"`
	MongodbatlasProcessOplogRate                          MetricConfig `mapstructure:"mongodbatlas.process.oplog.rate"`
	MongodbatlasProcessOplogTime                          MetricConfig `mapstructure:"mongodbatlas.process.oplog.time"`
	MongodbatlasProcessPageFaults                         MetricConfig `mapstructure:"mongodbatlas.process.page_faults"`
	MongodbatlasProcessRestarts                           MetricConfig `mapstructure:"mongodbatlas.process.restarts"`
	MongodbatlasProcessTickets                            MetricConfig `mapstructure:"mongodbatlas.process.tickets"`
	MongodbatlasSystemCPUNormalizedUsageAverage           MetricConfig `mapstructure:"mongodbatlas.system.cpu.normalized.usage.average"`
	MongodbatlasSystemCPUNormalizedUsageMax               MetricConfig `mapstructure:"mongodbatlas.system.cpu.normalized.usage.max"`
	MongodbatlasSystemCPUUsageAverage                     MetricConfig `mapstructure:"mongodbatlas.system.cpu.usage.average"`
	MongodbatlasSystemCPUUsageMax                         MetricConfig `mapstructure:"mongodbatlas.system.cpu.usage.max"`
	MongodbatlasSystemFtsCPUNormalizedUsage               MetricConfig `mapstructure:"mongodbatlas.system.fts.cpu.normalized.usage"`
	MongodbatlasSystemFtsCPUUsage                         MetricConfig `mapstructure:"mongodbatlas.system.fts.cpu.usage"`
	MongodbatlasSystemFtsDiskUsed                         MetricConfig `mapstructure:"mongodbatlas.system.fts.disk.used"`
	MongodbatlasSystemFtsMemoryUsage                      MetricConfig `mapstructure:"mongodbatlas.system.fts.memory.usage"`
	MongodbatlasSystemMemoryUsageAverage                  MetricConfig `mapstructure:"mongodbatlas.system.memory.usage.average"`
	MongodbatlasSystemMemoryUsageMax                      MetricConfig `mapstructure:"mongodbatlas.system.memory.usage.max"`
	MongodbatlasSystemNetworkIoAverage                    MetricConfig `mapstructure:"mongodbatlas.system.network.io.average"`
	MongodbatlasSystemNetworkIoMax                        MetricConfig `mapstructure:"mongodbatlas.system.network.io.max"`
	MongodbatlasSystemPagingIoAverage                     MetricConfig `mapstructure:"mongodbatlas.system.paging.io.average"`
	MongodbatlasSystemPagingIoMax                         MetricConfig `mapstructure:"mongodbatlas.system.paging.io.max"`
	MongodbatlasSystemPagingUsageAverage                  MetricConfig `mapstructure:"mongodbatlas.system.paging.usage.average"`
	MongodbatlasSystemPagingUsageMax                      MetricConfig `mapstructure:"mongodbatlas.system.paging.usage.max"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		MongodbatlasDbCounts: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDbSize: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionIopsAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionIopsMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionLatencyAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionLatencyMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionQueueDepth: MetricConfig{
			Enabled: false,
		},
		MongodbatlasDiskPartitionSpaceAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionSpaceMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionThroughput: MetricConfig{
			Enabled: false,
		},
		MongodbatlasDiskPartitionUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionUsageMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionUtilizationAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasDiskPartitionUtilizationMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessAsserts: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessBackgroundFlush: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCacheIo: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCacheRatio: MetricConfig{
			Enabled: false,
		},
		MongodbatlasProcessCacheSize: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessConnections: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCPUChildrenNormalizedUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCPUChildrenNormalizedUsageMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCPUChildrenUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCPUChildrenUsageMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCPUNormalizedUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCPUNormalizedUsageMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCPUUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCPUUsageMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessCursors: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessDbDocumentRate: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessDbOperationsRate: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessDbOperationsTime: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessDbQueryExecutorScanned: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessDbQueryTargetingScannedPerReturned: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessDbStorage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessGlobalLock: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessIndexBtreeMissRatio: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessIndexCounters: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessJournalingCommits: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessJournalingDataFiles: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessJournalingWritten: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessMemoryUsage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessNetworkIo: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessNetworkRequests: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessOplogRate: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessOplogTime: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessPageFaults: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessRestarts: MetricConfig{
			Enabled: true,
		},
		MongodbatlasProcessTickets: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemCPUNormalizedUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemCPUNormalizedUsageMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemCPUUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemCPUUsageMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemFtsCPUNormalizedUsage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemFtsCPUUsage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemFtsDiskUsed: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemFtsMemoryUsage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemMemoryUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemMemoryUsageMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemNetworkIoAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemNetworkIoMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemPagingIoAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemPagingIoMax: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemPagingUsageAverage: MetricConfig{
			Enabled: true,
		},
		MongodbatlasSystemPagingUsageMax: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for mongodbatlas resource attributes.
type ResourceAttributesConfig struct {
	MongodbAtlasClusterName     ResourceAttributeConfig `mapstructure:"mongodb_atlas.cluster.name"`
	MongodbAtlasDbName          ResourceAttributeConfig `mapstructure:"mongodb_atlas.db.name"`
	MongodbAtlasDiskPartition   ResourceAttributeConfig `mapstructure:"mongodb_atlas.disk.partition"`
	MongodbAtlasHostName        ResourceAttributeConfig `mapstructure:"mongodb_atlas.host.name"`
	MongodbAtlasOrgName         ResourceAttributeConfig `mapstructure:"mongodb_atlas.org_name"`
	MongodbAtlasProcessID       ResourceAttributeConfig `mapstructure:"mongodb_atlas.process.id"`
	MongodbAtlasProcessPort     ResourceAttributeConfig `mapstructure:"mongodb_atlas.process.port"`
	MongodbAtlasProcessTypeName ResourceAttributeConfig `mapstructure:"mongodb_atlas.process.type_name"`
	MongodbAtlasProjectID       ResourceAttributeConfig `mapstructure:"mongodb_atlas.project.id"`
	MongodbAtlasProjectName     ResourceAttributeConfig `mapstructure:"mongodb_atlas.project.name"`
	MongodbAtlasProviderName    ResourceAttributeConfig `mapstructure:"mongodb_atlas.provider.name"`
	MongodbAtlasRegionName      ResourceAttributeConfig `mapstructure:"mongodb_atlas.region.name"`
	MongodbAtlasUserAlias       ResourceAttributeConfig `mapstructure:"mongodb_atlas.user.alias"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		MongodbAtlasClusterName: ResourceAttributeConfig{
			Enabled: false,
		},
		MongodbAtlasDbName: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasDiskPartition: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasHostName: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasOrgName: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasProcessID: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasProcessPort: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasProcessTypeName: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasProjectID: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasProjectName: ResourceAttributeConfig{
			Enabled: true,
		},
		MongodbAtlasProviderName: ResourceAttributeConfig{
			Enabled: false,
		},
		MongodbAtlasRegionName: ResourceAttributeConfig{
			Enabled: false,
		},
		MongodbAtlasUserAlias: ResourceAttributeConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for mongodbatlas metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
