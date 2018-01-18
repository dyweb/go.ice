package sqlite

import (
	dlog "github.com/dyweb/gommon/log"
	"github.com/at15/go.ice/ice/db/drivers"
)

type Driver struct {
	log *dlog.Logger
}

func New() *Driver {
	d := &Driver{}
	d.log = dlog.NewStructLogger(log, d)
	return d
}

func (d *Driver) DefaultConfig() drivers.DefaultConfig {
	return defaultConfig
}
