package jaeger

import (
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	jg "github.com/uber/jaeger-client-go"
	jgconfig "github.com/uber/jaeger-client-go/config"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"

	"github.com/at15/go.ice/ice/config"
	"github.com/at15/go.ice/ice/tracing"
)

var _ tracing.Adapter = (*Adapter)(nil)

type Adapter struct {
	tracer opentracing.Tracer
	closer io.Closer
}

var _ jg.Logger = (*logger)(nil)

// wrapper to implement the strange jagger logger interface ...
type logger struct {
	service string
}

func newLogger(service string) *logger {
	return &logger{
		service: service,
	}
}

func (l *logger) Error(msg string) {
	log.ErrorF(msg, dlog.Fields{
		dlog.Str("svc", l.service),
		dlog.Str("trc", "jaeger"),
	})
}

func (l *logger) Infof(msg string, args ...interface{}) {
	log.ErrorF(fmt.Sprintf(msg, args...), dlog.Fields{
		dlog.Str("svc", l.service),
		dlog.Str("trc", "jaeger"),
	})
}

func (a *Adapter) NewTracer(service string, cfg config.TracingConfig) (opentracing.Tracer, error) {
	c := jgconfig.Configuration{
		Sampler: &jgconfig.SamplerConfig{
			Type:  cfg.Sampler.Type,
			Param: cfg.Sampler.Param,
		},
		Reporter: &jgconfig.ReporterConfig{
			LogSpans:            cfg.Reporter.LogSpans,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "localhost:6831",
		},
	}
	// TODO: Observer can be registered with the Tracer to receive notifications about new Spans.
	tracer, closer, err := c.New(service, jgconfig.Logger(newLogger(service)))
	if err != nil {
		return nil, errors.Wrap(err, "can't create jaeger tracer")
	}
	a.tracer = tracer
	a.closer = closer
	return tracer, nil
}

func (a *Adapter) Close() error {
	if err := a.closer.Close(); err != nil {
		// TODO: I think jaeger is using pkg/errors as well
		return errors.Wrap(err, "can't close jaeger tracer")
	}
	return nil
}
