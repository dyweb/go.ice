package jaeger

import (
	"github.com/at15/go.ice/ice/tracing"
	"github.com/at15/go.ice/ice/util/logutil"
)

const adapterName = "jaeger"

var log = logutil.NewPackageLogger()

func init() {
	tracing.RegisterAdapterFactory(adapterName, func() tracing.Adapter {
		return New()
	})
}
