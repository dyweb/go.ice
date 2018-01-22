package mysql

import (
	"github.com/at15/go.ice/ice/db"
	"github.com/at15/go.ice/ice/util/logutil"

	_ "github.com/go-sql-driver/mysql"
)

const adapterName = "mysql"
const driverName = "mysql"

var log = logutil.NewPackageLogger()

func init() {
	db.RegisterAdapterFactory(adapterName, func() db.Adapter {
		return New()
	})
}
