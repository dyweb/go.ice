package tracing

import (
	"github.com/at15/go.ice/ice/config"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type Manager struct {
	config  config.TracingConfig
	adapter Adapter
}

func NewManager(config config.TracingConfig) (*Manager, error) {
	adapter, err := GetAdapter(config.Adapter)
	if err != nil {
		return nil, errors.WithMessage(err, "unknown adapter "+config.Adapter)
	}
	return &Manager{
		adapter: adapter,
	}, nil
}

func (mgr *Manager) Tracer(service string) (opentracing.Tracer, error) {
	return mgr.adapter.NewTracer(service, mgr.config)
}
