// Code generated by gommon from db/drivers/postgres/gommon.yml DO NOT EDIT.
package postgres
import dlog "github.com/dyweb/gommon/log"

func (d *Driver) SetLogger(logger *dlog.Logger) {
	d.log = logger
}

func (d *Driver) GetLogger() *dlog.Logger {
	return d.log
}

func (d *Driver) LoggerIdentity(justCallMe func() *dlog.Identity) *dlog.Identity {
	return justCallMe()
}