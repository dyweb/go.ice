package postgres

import (
	"database/sql"
	"fmt"

	"github.com/at15/go.ice/ice/config"
	"github.com/at15/go.ice/ice/db"
	dlog "github.com/dyweb/gommon/log"
)

var _ db.Adapter = (*Adapter)(nil)

type Adapter struct {
	db  *sql.DB
	log *dlog.Logger
}

func New() *Adapter {
	a := &Adapter{}
	a.log = dlog.NewStructLogger(log, a)
	return a
}

func (a *Adapter) SetDB(db *sql.DB) {
	a.db = db
}

func (a *Adapter) GetDB() *sql.DB {
	if a.db == nil {
		a.log.Warn("db is nil pointer")
	}
	return a.db
}

func (a *Adapter) DriverName() string {
	return driverName
}

func (a *Adapter) Defaults() db.AdapterDefaults {
	return defaults
}

func (a *Adapter) FormatDSN(c config.DatabaseConfig) (string, error) {
	if c.DSN == "" {
		a.log.Debug("using DSN in config directly")
		return c.DSN, nil
	}
	user := c.User
	password := c.Password
	host := c.Host
	port := c.Port
	database := c.DBName
	sslmode := c.SSLMode
	if host == "" {
		host = "localhost"
	}
	if port <= 0 {
		port = defaults.Port()
	}
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d", user, password, host, port)
	if database != "" {
		dsn += " database=" + database
	}
	// TODO: what are valid sslmodes? should throw error here
	if sslmode != "" {
		dsn += " sslmode=" + sslmode
	}
	a.log.Debugf("format DSN based on config %s", dsn)
	return dsn, nil
}
