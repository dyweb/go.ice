package db

import (
	"sort"
	"sync"

	"github.com/at15/go.ice/ice/config"
)

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
