// Package mysql wraps github.com/go-sql-driver/mysql
package mysql

import (
	"github.com/dyweb/go.ice/ice/db"
	"github.com/dyweb/go.ice/ice/util/logutil"

	_ "github.com/go-sql-driver/mysql"
)

const adapterName = "mysql"
const driverName = "mysql"

var log, _ = logutil.NewPackageLoggerAndRegistry()

func init() {
	db.RegisterAdapterFactory(adapterName, func() db.Adapter {
		return New()
	})
}
