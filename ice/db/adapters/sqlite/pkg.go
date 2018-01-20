package sqlite

import (
	"github.com/at15/go.ice/ice/db"
	"github.com/at15/go.ice/ice/util/logutil"

	_ "github.com/mattn/go-sqlite3" // nameless import to register driver
)

const adapterName = "sqlite"
const driverName = "sqlite3"

var log = logutil.NewPackageLogger()

func init() {
	db.RegisterAdapterFactory(adapterName, func() db.Adapter {
		return New()
	})
}
