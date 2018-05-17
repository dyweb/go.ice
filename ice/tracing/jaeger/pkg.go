package jaeger

import (
	"github.com/dyweb/go.ice/ice/tracing"
	"github.com/dyweb/go.ice/ice/util/logutil"
)

const adapterName = "jaeger"

var log = logutil.NewPackageLogger()

func init() {
	tracing.RegisterAdapterFactory(adapterName, func() tracing.Adapter {
		return New()
	})
}
