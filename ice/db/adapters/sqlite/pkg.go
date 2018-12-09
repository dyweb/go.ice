// Package sqlite wraps github.com/mattn/go-sqlite3
package sqlite

import (
	"github.com/dyweb/go.ice/ice/db"
	"github.com/dyweb/go.ice/ice/util/logutil"

	_ "github.com/mattn/go-sqlite3" // nameless import to register driver
)

const adapterName = "sqlite"
const driverName = "sqlite3"

var log, _ = logutil.NewPackageLoggerAndRegistry()

func init() {
	db.RegisterAdapterFactory(adapterName, func() db.Adapter {
		return New()
	})
}
