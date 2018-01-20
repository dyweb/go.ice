package db

import (
	"database/sql"
	"fmt"

	"github.com/at15/go.ice/ice/config"
	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"
	"sync"
)

// TODO: future
// - each service should register which table it is using in manager, so it can print out the relationship
type Manager struct {
	mu     sync.RWMutex
	config config.DatabaseManagerConfig
	log    *dlog.Logger
	dbs    map[string]Adapter
}

func NewManager(config config.DatabaseManagerConfig) *Manager {
	m := &Manager{
		config: config,
		dbs:    make(map[string]Adapter, 1),
	}
	m.log = dlog.NewStructLogger(log, m)
	return m
}

func (mgr *Manager) Default() (Adapter, error) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	var (
		a     Adapter
		found = false
	)
	dbName := mgr.config.Default
	if a, found = mgr.dbs[dbName]; found {
		return a, nil
	}
	var c *config.DatabaseConfig = nil
	for _, d := range mgr.config.Databases {
		if d.Name == dbName {
			c = &d
			break
		}
	}
	if c == nil {
		return nil, errors.Errorf("default database %s is not configured", dbName)
	}
	var (
		err error
		dsn string
		db  *sql.DB
	)
	adapterName := c.Adapter
	if a, err = GetAdapter(adapterName); err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("can't get %s adapter for database %s", adapterName, dbName))
	}
	if dsn, err = a.FormatDSN(*c); err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("can't use %s adapter to format dsn for database %s", adapterName, dbName))
	}
	// NOTE: sql.Open does not make connection, so it won't throw error if remote db server is not ready
	if db, err = sql.Open(a.DriverName(), dsn); err != nil {
		return nil, errors.WithMessage(err, "can't open database handle")
	}
	a.SetDB(db)
	mgr.dbs[dbName] = a
	return a, nil
}

func (mgr *Manager) PrintConfig() {
	if mgr == nil {
		log.Warn("Manager is nil pointer")
		return
	}
	fmt.Printf("default %s\n", mgr.config.Default)
	fmt.Printf("enabled %s\n", mgr.config.Enabled)
	for i, c := range mgr.config.Databases {
		fmt.Printf("%d %s\n", i, c.String())
	}
}

func (mgr *Manager) RegisteredDrivers() []string {
	return Drivers()
}

func (mgr *Manager) RegisteredAdapters() []string {
	return Adapters()
}

func Drivers() []string {
	return sql.Drivers()
}
