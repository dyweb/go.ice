package tracing

import "github.com/opentracing/opentracing-go"

type Manager struct {
	tracer opentracing.Tracer
}
