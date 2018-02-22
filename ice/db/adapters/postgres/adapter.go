package postgres

import (
	"fmt"

	"github.com/at15/go.ice/ice/config"
	"github.com/at15/go.ice/ice/db"
	dlog "github.com/dyweb/gommon/log"
)

var _ db.Adapter = (*Adapter)(nil)

type Adapter struct {
	log *dlog.Logger
}

func New() *Adapter {
	a := &Adapter{}
	a.log = dlog.NewStructLogger(log, a)
	return a
}

func (a *Adapter) DriverName() string {
	return driverName
}

func (a *Adapter) Defaults() db.AdapterDefaults {
	return defaults
}

func (a *Adapter) FormatDSN(c config.DatabaseConfig) (string, error) {
	if c.DSN != "" {
		a.log.Debugf("using DSN in config directly %s", c.DSN)
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

func (a *Adapter) CanCreateDatabase() bool {
	return true
}
