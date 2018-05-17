package tracing

import (
	"sort"
	"sync"

	"github.com/dyweb/gommon/errors"
	"github.com/opentracing/opentracing-go"

	"github.com/dyweb/go.ice/ice/config"
)

var (
	adaptersMu        sync.RWMutex
	adaptersFactories = make(map[string]AdapterFactory)
)

type AdapterFactory func() Adapter

type Adapter interface {
	NewTracer(service string, cfg config.TracingConfig) (opentracing.Tracer, error)
	Close() error
}

// TODO: future, adapter factory functions should allow passing config
func GetAdapter(name string) (Adapter, error) {
	adaptersMu.RLock()
	defer adaptersMu.RUnlock()
	if f, ok := adaptersFactories[name]; !ok {
		return nil, errors.Errorf("adapter %s is not registered", name)
	} else {
		return f(), nil
	}
}

func RegisterAdapterFactory(name string, factory AdapterFactory) {
	adaptersMu.Lock()
	defer adaptersMu.Unlock()
	if _, dup := adaptersFactories[name]; dup {
		log.Panicf("RegisterAdapterFactory %s is called twice", name)
	}
	adaptersFactories[name] = factory
}

func Adapters() []string {
	adaptersMu.RLock()
	defer adaptersMu.RUnlock()
	var list []string
	for name := range adaptersFactories {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}
