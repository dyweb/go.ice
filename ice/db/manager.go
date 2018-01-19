package db

import (
	"database/sql"
	"fmt"

	dlog "github.com/dyweb/gommon/log"
	"github.com/at15/go.ice/ice/config"
)

// TODO: future
// - each service should register which table it is using in manager, so it can print out the relationship
type Manager struct {
	config config.DatabaseManagerConfig
	log    *dlog.Logger
}

func NewManager(config config.DatabaseManagerConfig) *Manager {
	m :=  &Manager{
		config: config,
	}
	m.log = dlog.NewStructLogger(log, m)
	return m
}

func (mgr *Manager) PrintConfig() {
	if mgr == nil {
		// TODO: warn using logger
		fmt.Println("empty mgr")
	}
	fmt.Printf("default %s\n", mgr.config.Default)
	fmt.Printf("enabled %s\n", mgr.config.Enabled)
	for i, c := range mgr.config.Databases {
		fmt.Printf("%d %s\n", i, c.String())
	}
}

func (mgr *Manager) NativeDrivers() []string {
	return NativeDrivers()
}

func NativeDrivers() []string {
	return sql.Drivers()
}
