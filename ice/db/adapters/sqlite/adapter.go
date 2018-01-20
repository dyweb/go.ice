package sqlite

import (
	"github.com/pkg/errors"

	"github.com/at15/go.ice/ice/config"
	"github.com/at15/go.ice/ice/db"
	dlog "github.com/dyweb/gommon/log"
	"database/sql"
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
		return "", errors.New("sqlite3 must specify file path as dsn in config")
	}
	return c.DSN, nil
}
