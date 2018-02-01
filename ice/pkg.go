// Package ice TODO: ... say something (I'm giving up on you?)
package ice

import "github.com/at15/go.ice/ice/tracing"

// TODO: do I really need the service package
type Service interface {
	SetTracer(tracer tracing.Tracer)
	Tracer() tracing.Tracer
}
