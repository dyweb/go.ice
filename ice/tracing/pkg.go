// Package tracing TODO: should do something ...
package tracing

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	jgconfig "github.com/uber/jaeger-client-go/config"
)

var tracer opentracing.Tracer
var closer io.Closer

// FIXME: hacky function to play with tracing libraries
// https://github.com/jaegertracing/jaeger/blob/master/examples/hotrod/pkg/tracing/init.go#L31
func configTracer() error {
	tcfg := jgconfig.Configuration{
		Sampler: &jgconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jgconfig.ReporterConfig{
			LogSpans: false, // TODO: when true, enables LoggingReporter that runs in parallel with the main reporter
			// and logs all submitted spans. Main Configuration.Logger must be initialized in the code
			// for this option to have any effect.
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "localhost:6831",
		},
	}
	// TODO: a better way to use gommon/log, current tree level hierarchy may not be enough ...
	// TODO: the jaeger.Logger interface is so strange, Error(string) instead of Error(string, args ...interface{})
	// jgconfig.Logger(log)
	// TODO: Observer can be registered with the Tracer to receive notifications about new Spans.
	var err error
	tracer, closer, err = tcfg.New("service-a")
	if err != nil {
		return errors.Wrap(err, "can't create jaeger tracer")
	}
	return nil
}
