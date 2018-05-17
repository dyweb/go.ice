package jaeger

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	jg "github.com/uber/jaeger-client-go"
	jgconfig "github.com/uber/jaeger-client-go/config"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	"github.com/dyweb/go.ice/ice/config"
	"github.com/dyweb/go.ice/ice/tracing"
)

var _ tracing.Adapter = (*Adapter)(nil)

type Adapter struct {
	tracers map[string]opentracing.Tracer
	closers map[string]io.Closer
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
	log.InfoF(fmt.Sprintf(msg, args...), dlog.Fields{
		dlog.Str("svc", l.service),
		dlog.Str("trc", "jaeger"),
	})
}

func New() *Adapter {
	a := &Adapter{
		tracers: make(map[string]opentracing.Tracer, 5),
		closers: make(map[string]io.Closer, 5),
	}
	// TODO: gommon struct logger
	return a
}

func (a *Adapter) NewTracer(service string, cfg config.TracingConfig) (opentracing.Tracer, error) {
	if tracer, exists := a.tracers[service]; exists {
		log.Warnf("reuse existing tracers for service %s", service)
		return tracer, nil
	}
	log.Debugf("jaeger tracer sampler type %s param %f", cfg.Sampler.Type, cfg.Sampler.Param)
	log.Debugf("jaeger tracer reporter agent %s logspan %v", cfg.Reporter.LocalAgentHostPort, cfg.Reporter.LogSpans)
	c := jgconfig.Configuration{
		Sampler: &jgconfig.SamplerConfig{
			Type:  cfg.Sampler.Type,
			Param: cfg.Sampler.Param,
		},
		Reporter: &jgconfig.ReporterConfig{
			LogSpans: cfg.Reporter.LogSpans,
			// TODO: allow config interval in config
			//BufferFlushInterval: 10 * time.Millisecond,
			LocalAgentHostPort: cfg.Reporter.LocalAgentHostPort,
		},
	}
	// TODO: Observer can be registered with the Tracer to receive notifications about new Spans.
	tracer, closer, err := c.New(service, jgconfig.Logger(newLogger(service)))
	if err != nil {
		return nil, errors.Wrap(err, "can't create jaeger tracer")
	}
	a.tracers[service] = tracer
	a.closers[service] = closer
	return tracer, nil
}

func (a *Adapter) Close() error {
	// TODO: might need a error group instead of just return the last one
	var lastError error
	for service, closer := range a.closers {
		if err := closer.Close(); err != nil {
			lastError = errors.Wrap(err, "can't close jaeger tracer for service "+service)
			log.Warn(lastError)
		}
	}
	return lastError
}
