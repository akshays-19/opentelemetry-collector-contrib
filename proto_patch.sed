s+github.com/open-telemetry/opentelemetry-proto/gen/go/+go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/+g

s+package opentelemetry.proto.\(.*\).v1;+package opentelemetry.proto.\1.v1;\
\
import "gogoproto/gogo.proto";+g

s+bytes trace_id = \(.*\);+bytes trace_id = \1\
  [\
  // Use custom TraceId data type for this field.\
  (gogoproto.nullable) = false,\
  (gogoproto.customtype) = "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/common/v1.TraceID"\
  ];+g

s+bytes \(.*span_id\) = \(.*\);+bytes \1 = \2\
  [\
  // Use custom SpanId data type for this field.\
  (gogoproto.nullable) = false,\
  (gogoproto.customtype) = "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/common/v1.SpanID"\
  ];+g
