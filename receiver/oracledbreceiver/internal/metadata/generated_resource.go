// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetHostName sets provided value as "host.name" attribute.
func (rb *ResourceBuilder) SetHostName(val string) {
	if rb.config.HostName.Enabled {
		rb.res.Attributes().PutStr("host.name", val)
	}
}

// SetOracledbInstanceName sets provided value as "oracledb.instance.name" attribute.
func (rb *ResourceBuilder) SetOracledbInstanceName(val string) {
	if rb.config.OracledbInstanceName.Enabled {
		rb.res.Attributes().PutStr("oracledb.instance.name", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
