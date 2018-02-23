package mysql

import (
	"fmt"
	"strings"

	"github.com/at15/go.ice/ice/config"
	"github.com/at15/go.ice/ice/db"
	dlog "github.com/dyweb/gommon/log"
	"github.com/go-sql-driver/mysql"
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
	mc := mysql.NewConfig()
	mc.User = c.User
	mc.Passwd = c.Password
	mc.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
	mc.DBName = c.DBName
	dsn := mc.FormatDSN()
	a.log.Debugf("format DSN based on config %s", dsn)
	return dsn, nil
}

func (a *Adapter) CanCreateDatabase() bool {
	return true
}

// based on https://github.com/Masterminds/squirrel/blob/v1/placeholder.go
func (a *Adapter) Placeholders(count int) string {
	if count < 1 {
		return ""
	}
	return strings.Repeat(",?", count)[1:]
}
