package sqlite

import (
	dlog "github.com/dyweb/gommon/log"
	"github.com/at15/go.ice/ice/db"
)

var _ db.Driver = (*Driver)(nil)

type Driver struct {
	log *dlog.Logger
}

func New() *Driver {
	d := &Driver{}
	d.log = dlog.NewStructLogger(log, d)
	return d
}

func (d *Driver) Defaults() db.DriverDefaults {
	return defaults
}
