package postgres

import (
	"fmt"
	"strconv"

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

// based on https://github.com/Masterminds/squirrel/blob/v1/placeholder.go
func (a *Adapter) Placeholders(count int) string {
	if count < 1 {
		return ""
	}
	buf := make([]byte, 0, count*3)
	for i := 1; i <= count; i++ {
		buf = append(buf, ',', '$')
		buf = strconv.AppendInt(buf, int64(i), 10)
	}
	return string(buf[1:])
}
