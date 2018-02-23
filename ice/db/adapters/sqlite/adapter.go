package sqlite

import (
	"github.com/pkg/errors"
	"strings"

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
	if c.DSN == "" {
		return "", errors.New("sqlite3 must specify file path as dsn in config")
	}
	return c.DSN, nil
}

func (a *Adapter) CanCreateDatabase() bool {
	return false
}

// based on https://github.com/Masterminds/squirrel/blob/v1/placeholder.go
func (a *Adapter) Placeholders(count int) string {
	if count < 1 {
		return ""
	}
	return strings.Repeat(",?", count)[1:]
}
