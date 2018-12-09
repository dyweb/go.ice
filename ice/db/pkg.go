// Package db defines interface for using RDBMS
package db // import "github.com/dyweb/go.ice/ice/db"

import (
	"database/sql"
	"github.com/dyweb/go.ice/ice/util/logutil"
)

var log, _ = logutil.NewPackageLoggerAndRegistry()

func Drivers() []string {
	return sql.Drivers()
}
