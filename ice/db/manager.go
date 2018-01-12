package db

import (
	"database/sql"
)

// TODO: future
// - each service should register which table it is using in manager, so it can print out the relationship
type Manager struct {
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
