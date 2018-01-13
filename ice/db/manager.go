package db

import (
	"database/sql"
	"github.com/at15/go.ice/ice/config"
	"fmt"
)

// TODO: future
// - each service should register which table it is using in manager, so it can print out the relationship
// - generate logger interface function
type Manager struct {
	config config.DatabaseManagerConfig
}

func NewManager(config config.DatabaseManagerConfig) *Manager {
	return &Manager{
		config: config,
	}
}

func (mgr *Manager) PrintConfig() {
	fmt.Printf("default %s\n", mgr.config.Default)
	fmt.Printf("enabled %s\n", mgr.config.Enabled)
	for i, c := range mgr.config.Databases {
		fmt.Printf("%d %s\n", i, c.String())
	}
}

func (mgr *Manager) HasDriver(driver string) bool {
	return HasDriver(driver)
}

func (mgr *Manager) Drivers() []string {
	return Drivers()
}

func HasDriver(driver string) bool {
	// TODO: tolerate common names?
	for _, d := range Drivers() {
		if d == driver {
			return true
		}
	}
	return false
}

func Drivers() []string {
	return sql.Drivers()
}
