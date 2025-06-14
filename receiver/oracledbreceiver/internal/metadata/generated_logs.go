// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/filter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/otel/trace"
)

type eventDbServerQuerySample struct {
	data   plog.LogRecordSlice // data buffer for generated log records.
	config EventConfig         // event config provided by user.
}

func (e *eventDbServerQuerySample) recordEvent(ctx context.Context, timestamp pcommon.Timestamp, dbQueryTextAttributeValue string, dbSystemNameAttributeValue string, dbQueryPlanHashValueAttributeValue string, dbMachineAttributeValue string, dbQueryIDAttributeValue string, oracledbQueryChildNumberAttributeValue string, dbQuerySessionIDAttributeValue string, dbQuerySerialNumberAttributeValue string, dbQueryProcessAttributeValue string, oracledbUsernameAttributeValue string, oracledbSchemaNameAttributeValue string, oracledbQueryProgramAttributeValue string, oracledbQueryModuleAttributeValue string, oracledbQueryStatusAttributeValue string, oracledbQueryStateAttributeValue string, oracledbQueryWaitClassAttributeValue string, oracledbQueryEventAttributeValue string, oracledbQueryObjectNameAttributeValue string, oracledbQueryObjectTypeAttributeValue string, oracledbQueryOsUserAttributeValue string, oracledbQueryDurationAttributeValue float64) {
	if !e.config.Enabled {
		return
	}
	lr := e.data.AppendEmpty()
	lr.SetEventName("db.server.query_sample")
	lr.SetTimestamp(timestamp)

	if span := trace.SpanContextFromContext(ctx); span.IsValid() {
		lr.SetTraceID(pcommon.TraceID(span.TraceID()))
		lr.SetSpanID(pcommon.SpanID(span.SpanID()))
	}

	lr.Attributes().PutStr("db.query.text", dbQueryTextAttributeValue)
	lr.Attributes().PutStr("db.system.name", dbSystemNameAttributeValue)
	lr.Attributes().PutStr("db.query.plan_hash_value", dbQueryPlanHashValueAttributeValue)
	lr.Attributes().PutStr("db.machine", dbMachineAttributeValue)
	lr.Attributes().PutStr("db.query.id", dbQueryIDAttributeValue)
	lr.Attributes().PutStr("oracledb.query.child_number", oracledbQueryChildNumberAttributeValue)
	lr.Attributes().PutStr("db.query.session_id", dbQuerySessionIDAttributeValue)
	lr.Attributes().PutStr("db.query.serial_number", dbQuerySerialNumberAttributeValue)
	lr.Attributes().PutStr("db.query.process", dbQueryProcessAttributeValue)
	lr.Attributes().PutStr("oracledb.username", oracledbUsernameAttributeValue)
	lr.Attributes().PutStr("oracledb.schema_name", oracledbSchemaNameAttributeValue)
	lr.Attributes().PutStr("oracledb.query.program", oracledbQueryProgramAttributeValue)
	lr.Attributes().PutStr("oracledb.query.module", oracledbQueryModuleAttributeValue)
	lr.Attributes().PutStr("oracledb.query.status", oracledbQueryStatusAttributeValue)
	lr.Attributes().PutStr("oracledb.query.state", oracledbQueryStateAttributeValue)
	lr.Attributes().PutStr("oracledb.query.wait_class", oracledbQueryWaitClassAttributeValue)
	lr.Attributes().PutStr("oracledb.query.event", oracledbQueryEventAttributeValue)
	lr.Attributes().PutStr("oracledb.query.object_name", oracledbQueryObjectNameAttributeValue)
	lr.Attributes().PutStr("oracledb.query.object_type", oracledbQueryObjectTypeAttributeValue)
	lr.Attributes().PutStr("oracledb.query.os_user", oracledbQueryOsUserAttributeValue)
	lr.Attributes().PutDouble("oracledb.query.duration", oracledbQueryDurationAttributeValue)
}

// emit appends recorded event data to a events slice and prepares it for recording another set of log records.
func (e *eventDbServerQuerySample) emit(lrs plog.LogRecordSlice) {
	if e.config.Enabled && e.data.Len() > 0 {
		e.data.MoveAndAppendTo(lrs)
	}
}

func newEventDbServerQuerySample(cfg EventConfig) eventDbServerQuerySample {
	e := eventDbServerQuerySample{config: cfg}
	if cfg.Enabled {
		e.data = plog.NewLogRecordSlice()
	}
	return e
}

type eventDbServerTopQuery struct {
	data   plog.LogRecordSlice // data buffer for generated log records.
	config EventConfig         // event config provided by user.
}

func (e *eventDbServerTopQuery) recordEvent(ctx context.Context, timestamp pcommon.Timestamp, dbQueryTextAttributeValue string, oracledbQueryPlanAttributeValue string, oracledbQuerySQLIDAttributeValue string, oracledbQueryChildNumberAttributeValue string, oracledbQueryApplicationWaitTimeAttributeValue float64, oracledbQueryBufferGetsAttributeValue int64, oracledbQueryClusterWaitTimeAttributeValue float64, oracledbQueryConcurrencyWaitTimeAttributeValue float64, oracledbQueryCPUTimeAttributeValue float64, oracledbQueryDirectReadsAttributeValue int64, oracledbQueryDirectWritesAttributeValue int64, oracledbQueryDiskReadsAttributeValue int64, oracledbQueryElapsedTimeAttributeValue float64, oracledbQueryExecutionsAttributeValue int64, oracledbQueryPhysicalReadBytesAttributeValue int64, oracledbQueryPhysicalReadRequestsAttributeValue int64, oracledbQueryPhysicalWriteBytesAttributeValue int64, oracledbQueryPhysicalWriteRequestsAttributeValue int64, oracledbQueryRowsProcessedAttributeValue int64, oracledbQueryUserIoWaitTimeAttributeValue float64, dbServerNameAttributeValue string) {
	if !e.config.Enabled {
		return
	}
	lr := e.data.AppendEmpty()
	lr.SetEventName("db.server.top_query")
	lr.SetTimestamp(timestamp)

	if span := trace.SpanContextFromContext(ctx); span.IsValid() {
		lr.SetTraceID(pcommon.TraceID(span.TraceID()))
		lr.SetSpanID(pcommon.SpanID(span.SpanID()))
	}

	lr.Attributes().PutStr("db.query.text", dbQueryTextAttributeValue)
	lr.Attributes().PutStr("oracledb.query_plan", oracledbQueryPlanAttributeValue)
	lr.Attributes().PutStr("oracledb.query.sql_id", oracledbQuerySQLIDAttributeValue)
	lr.Attributes().PutStr("oracledb.query.child_number", oracledbQueryChildNumberAttributeValue)
	lr.Attributes().PutDouble("oracledb.query.application_wait_time", oracledbQueryApplicationWaitTimeAttributeValue)
	lr.Attributes().PutInt("oracledb.query.buffer_gets", oracledbQueryBufferGetsAttributeValue)
	lr.Attributes().PutDouble("oracledb.query.cluster_wait_time", oracledbQueryClusterWaitTimeAttributeValue)
	lr.Attributes().PutDouble("oracledb.query.concurrency_wait_time", oracledbQueryConcurrencyWaitTimeAttributeValue)
	lr.Attributes().PutDouble("oracledb.query.cpu_time", oracledbQueryCPUTimeAttributeValue)
	lr.Attributes().PutInt("oracledb.query.direct_reads", oracledbQueryDirectReadsAttributeValue)
	lr.Attributes().PutInt("oracledb.query.direct_writes", oracledbQueryDirectWritesAttributeValue)
	lr.Attributes().PutInt("oracledb.query.disk_reads", oracledbQueryDiskReadsAttributeValue)
	lr.Attributes().PutDouble("oracledb.query.elapsed_time", oracledbQueryElapsedTimeAttributeValue)
	lr.Attributes().PutInt("oracledb.query.executions", oracledbQueryExecutionsAttributeValue)
	lr.Attributes().PutInt("oracledb.query.physical_read_bytes", oracledbQueryPhysicalReadBytesAttributeValue)
	lr.Attributes().PutInt("oracledb.query.physical_read_requests", oracledbQueryPhysicalReadRequestsAttributeValue)
	lr.Attributes().PutInt("oracledb.query.physical_write_bytes", oracledbQueryPhysicalWriteBytesAttributeValue)
	lr.Attributes().PutInt("oracledb.query.physical_write_requests", oracledbQueryPhysicalWriteRequestsAttributeValue)
	lr.Attributes().PutInt("oracledb.query.rows_processed", oracledbQueryRowsProcessedAttributeValue)
	lr.Attributes().PutDouble("oracledb.query.user_io_wait_time", oracledbQueryUserIoWaitTimeAttributeValue)
	lr.Attributes().PutStr("db.server.name", dbServerNameAttributeValue)
}

// emit appends recorded event data to a events slice and prepares it for recording another set of log records.
func (e *eventDbServerTopQuery) emit(lrs plog.LogRecordSlice) {
	if e.config.Enabled && e.data.Len() > 0 {
		e.data.MoveAndAppendTo(lrs)
	}
}

func newEventDbServerTopQuery(cfg EventConfig) eventDbServerTopQuery {
	e := eventDbServerTopQuery{config: cfg}
	if cfg.Enabled {
		e.data = plog.NewLogRecordSlice()
	}
	return e
}

// LogsBuilder provides an interface for scrapers to report logs while taking care of all the transformations
// required to produce log representation defined in metadata and user config.
type LogsBuilder struct {
	config                         LogsBuilderConfig // config of the logs builder.
	logsBuffer                     plog.Logs
	logRecordsBuffer               plog.LogRecordSlice
	buildInfo                      component.BuildInfo // contains version information.
	resourceAttributeIncludeFilter map[string]filter.Filter
	resourceAttributeExcludeFilter map[string]filter.Filter
	eventDbServerQuerySample       eventDbServerQuerySample
	eventDbServerTopQuery          eventDbServerTopQuery
}

// LogBuilderOption applies changes to default logs builder.
type LogBuilderOption interface {
	apply(*LogsBuilder)
}

func NewLogsBuilder(lbc LogsBuilderConfig, settings receiver.Settings) *LogsBuilder {
	lb := &LogsBuilder{
		config:                         lbc,
		logsBuffer:                     plog.NewLogs(),
		logRecordsBuffer:               plog.NewLogRecordSlice(),
		buildInfo:                      settings.BuildInfo,
		eventDbServerQuerySample:       newEventDbServerQuerySample(lbc.Events.DbServerQuerySample),
		eventDbServerTopQuery:          newEventDbServerTopQuery(lbc.Events.DbServerTopQuery),
		resourceAttributeIncludeFilter: make(map[string]filter.Filter),
		resourceAttributeExcludeFilter: make(map[string]filter.Filter),
	}
	if lbc.ResourceAttributes.HostName.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["host.name"] = filter.CreateFilter(lbc.ResourceAttributes.HostName.EventsInclude)
	}
	if lbc.ResourceAttributes.HostName.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["host.name"] = filter.CreateFilter(lbc.ResourceAttributes.HostName.EventsExclude)
	}
	if lbc.ResourceAttributes.OracledbInstanceName.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["oracledb.instance.name"] = filter.CreateFilter(lbc.ResourceAttributes.OracledbInstanceName.EventsInclude)
	}
	if lbc.ResourceAttributes.OracledbInstanceName.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["oracledb.instance.name"] = filter.CreateFilter(lbc.ResourceAttributes.OracledbInstanceName.EventsExclude)
	}

	return lb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted logs.
func (lb *LogsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(lb.config.ResourceAttributes)
}

// ResourceLogsOption applies changes to provided resource logs.
type ResourceLogsOption interface {
	apply(plog.ResourceLogs)
}

type resourceLogsOptionFunc func(plog.ResourceLogs)

func (rlof resourceLogsOptionFunc) apply(rl plog.ResourceLogs) {
	rlof(rl)
}

// WithLogsResource sets the provided resource on the emitted ResourceLogs.
// It's recommended to use ResourceBuilder to create the resource.
func WithLogsResource(res pcommon.Resource) ResourceLogsOption {
	return resourceLogsOptionFunc(func(rl plog.ResourceLogs) {
		res.CopyTo(rl.Resource())
	})
}

// AppendLogRecord adds a log record to the logs builder.
func (lb *LogsBuilder) AppendLogRecord(lr plog.LogRecord) {
	lr.MoveTo(lb.logRecordsBuffer.AppendEmpty())
}

// EmitForResource saves all the generated logs under a new resource and updates the internal state to be ready for
// recording another set of log records as part of another resource. This function can be helpful when one scraper
// needs to emit logs from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceLogsOption arguments.
func (lb *LogsBuilder) EmitForResource(options ...ResourceLogsOption) {
	rl := plog.NewResourceLogs()
	ils := rl.ScopeLogs().AppendEmpty()
	ils.Scope().SetName(ScopeName)
	ils.Scope().SetVersion(lb.buildInfo.Version)
	lb.eventDbServerQuerySample.emit(ils.LogRecords())
	lb.eventDbServerTopQuery.emit(ils.LogRecords())

	for _, op := range options {
		op.apply(rl)
	}

	if lb.logRecordsBuffer.Len() > 0 {
		lb.logRecordsBuffer.MoveAndAppendTo(ils.LogRecords())
		lb.logRecordsBuffer = plog.NewLogRecordSlice()
	}

	for attr, filter := range lb.resourceAttributeIncludeFilter {
		if val, ok := rl.Resource().Attributes().Get(attr); ok && !filter.Matches(val.AsString()) {
			return
		}
	}
	for attr, filter := range lb.resourceAttributeExcludeFilter {
		if val, ok := rl.Resource().Attributes().Get(attr); ok && filter.Matches(val.AsString()) {
			return
		}
	}

	if ils.LogRecords().Len() > 0 {
		rl.MoveTo(lb.logsBuffer.ResourceLogs().AppendEmpty())
	}
}

// Emit returns all the logs accumulated by the logs builder and updates the internal state to be ready for
// recording another set of logs. This function will be responsible for applying all the transformations required to
// produce logs representation defined in metadata and user config.
func (lb *LogsBuilder) Emit(options ...ResourceLogsOption) plog.Logs {
	lb.EmitForResource(options...)
	logs := lb.logsBuffer
	lb.logsBuffer = plog.NewLogs()
	return logs
}

// RecordDbServerQuerySampleEvent adds a log record of db.server.query_sample event.
func (lb *LogsBuilder) RecordDbServerQuerySampleEvent(ctx context.Context, timestamp pcommon.Timestamp, dbQueryTextAttributeValue string, dbSystemNameAttributeValue string, dbQueryPlanHashValueAttributeValue string, dbMachineAttributeValue string, dbQueryIDAttributeValue string, oracledbQueryChildNumberAttributeValue string, dbQuerySessionIDAttributeValue string, dbQuerySerialNumberAttributeValue string, dbQueryProcessAttributeValue string, oracledbUsernameAttributeValue string, oracledbSchemaNameAttributeValue string, oracledbQueryProgramAttributeValue string, oracledbQueryModuleAttributeValue string, oracledbQueryStatusAttributeValue string, oracledbQueryStateAttributeValue string, oracledbQueryWaitClassAttributeValue string, oracledbQueryEventAttributeValue string, oracledbQueryObjectNameAttributeValue string, oracledbQueryObjectTypeAttributeValue string, oracledbQueryOsUserAttributeValue string, oracledbQueryDurationAttributeValue float64) {
	lb.eventDbServerQuerySample.recordEvent(ctx, timestamp, dbQueryTextAttributeValue, dbSystemNameAttributeValue, dbQueryPlanHashValueAttributeValue, dbMachineAttributeValue, dbQueryIDAttributeValue, oracledbQueryChildNumberAttributeValue, dbQuerySessionIDAttributeValue, dbQuerySerialNumberAttributeValue, dbQueryProcessAttributeValue, oracledbUsernameAttributeValue, oracledbSchemaNameAttributeValue, oracledbQueryProgramAttributeValue, oracledbQueryModuleAttributeValue, oracledbQueryStatusAttributeValue, oracledbQueryStateAttributeValue, oracledbQueryWaitClassAttributeValue, oracledbQueryEventAttributeValue, oracledbQueryObjectNameAttributeValue, oracledbQueryObjectTypeAttributeValue, oracledbQueryOsUserAttributeValue, oracledbQueryDurationAttributeValue)
}

// RecordDbServerTopQueryEvent adds a log record of db.server.top_query event.
func (lb *LogsBuilder) RecordDbServerTopQueryEvent(ctx context.Context, timestamp pcommon.Timestamp, dbQueryTextAttributeValue string, oracledbQueryPlanAttributeValue string, oracledbQuerySQLIDAttributeValue string, oracledbQueryChildNumberAttributeValue string, oracledbQueryApplicationWaitTimeAttributeValue float64, oracledbQueryBufferGetsAttributeValue int64, oracledbQueryClusterWaitTimeAttributeValue float64, oracledbQueryConcurrencyWaitTimeAttributeValue float64, oracledbQueryCPUTimeAttributeValue float64, oracledbQueryDirectReadsAttributeValue int64, oracledbQueryDirectWritesAttributeValue int64, oracledbQueryDiskReadsAttributeValue int64, oracledbQueryElapsedTimeAttributeValue float64, oracledbQueryExecutionsAttributeValue int64, oracledbQueryPhysicalReadBytesAttributeValue int64, oracledbQueryPhysicalReadRequestsAttributeValue int64, oracledbQueryPhysicalWriteBytesAttributeValue int64, oracledbQueryPhysicalWriteRequestsAttributeValue int64, oracledbQueryRowsProcessedAttributeValue int64, oracledbQueryUserIoWaitTimeAttributeValue float64, dbServerNameAttributeValue string) {
	lb.eventDbServerTopQuery.recordEvent(ctx, timestamp, dbQueryTextAttributeValue, oracledbQueryPlanAttributeValue, oracledbQuerySQLIDAttributeValue, oracledbQueryChildNumberAttributeValue, oracledbQueryApplicationWaitTimeAttributeValue, oracledbQueryBufferGetsAttributeValue, oracledbQueryClusterWaitTimeAttributeValue, oracledbQueryConcurrencyWaitTimeAttributeValue, oracledbQueryCPUTimeAttributeValue, oracledbQueryDirectReadsAttributeValue, oracledbQueryDirectWritesAttributeValue, oracledbQueryDiskReadsAttributeValue, oracledbQueryElapsedTimeAttributeValue, oracledbQueryExecutionsAttributeValue, oracledbQueryPhysicalReadBytesAttributeValue, oracledbQueryPhysicalReadRequestsAttributeValue, oracledbQueryPhysicalWriteBytesAttributeValue, oracledbQueryPhysicalWriteRequestsAttributeValue, oracledbQueryRowsProcessedAttributeValue, oracledbQueryUserIoWaitTimeAttributeValue, dbServerNameAttributeValue)
}
