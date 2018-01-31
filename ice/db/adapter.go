package db

import (
	"sort"
	"sync"

	"github.com/at15/go.ice/ice/config"
	"github.com/pkg/errors"
)

// TODO: allow adapter to wrap common sql operation and trace it? or it should be done at wrapper level?
// TODO: might change to call this dialect? because it need to deal with things like replace $ with ?

var (
	adaptersMu        sync.RWMutex
	adaptersFactories = make(map[string]AdapterFactory)
)

type AdapterDefaults interface {
	DockerImage() string
	Port() int
	NativeShell() string
}

// Adapter is high level wrapper around underlying drivers
type Adapter interface {
	DriverName() string
	Defaults() AdapterDefaults
	FormatDSN(c config.DatabaseConfig) (string, error)
}

type AdapterFactory func() Adapter

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
