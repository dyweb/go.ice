// Package db defines interface for using RDBMS
package db

import (
	"database/sql"
	"github.com/at15/go.ice/ice/util/logutil"
)

var log = logutil.NewPackageLogger()

func Drivers() []string {
	return sql.Drivers()
}
