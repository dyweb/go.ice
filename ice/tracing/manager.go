package tracing

import (
	"github.com/dyweb/gommon/errors"
	"github.com/opentracing/opentracing-go"

	"github.com/dyweb/go.ice/ice/config"
)

type Manager struct {
	config  config.TracingConfig
	adapter Adapter
}

func NewManager(config config.TracingConfig) (*Manager, error) {
	adapter, err := GetAdapter(config.Adapter)
	if err != nil {
		return nil, errors.Wrap(err, "unknown adapter "+config.Adapter)
	}
	return &Manager{
		config:  config, // NOTE: took me a long time to find out where things went wrong ...
		adapter: adapter,
	}, nil
}

func (mgr *Manager) Tracer(service string) (opentracing.Tracer, error) {
	return mgr.adapter.NewTracer(service, mgr.config)
}
