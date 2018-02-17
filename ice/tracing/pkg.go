// Package tracing TODO: should do something ...
package tracing

import (
	"github.com/at15/go.ice/ice/config"
	"github.com/opentracing/opentracing-go"
)

type Adapter interface {
	NewTracer(service string, cfg config.TracingConfig) (opentracing.Tracer, error)
	Close() error
}
