package db

import (
	"sync"
	"sort"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]Driver)
)

type DriverDefaults interface {
	DockerImage() string
	Port() int
	NativeShell() string
}

// Driver is high level wrapper around underlying drivers
type Driver interface {
	Defaults() DriverDefaults
}

func RegisterDriver(name string, driver Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		log.Panicf("Register driver %s is nil", name)
	}
	if _, dup := drivers[name]; dup {
		log.Panicf("Register drivers %s is called twice", name)
	}
	drivers[name] = driver
}

func Drivers() []string {
	driversMu.RLock()
	defer driversMu.RUnlock()
	var list []string
	for name := range drivers {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}
