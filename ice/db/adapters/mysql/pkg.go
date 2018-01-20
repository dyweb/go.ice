package mysql

import (
	"github.com/at15/go.ice/ice/util/logutil"

	_ "github.com/go-sql-driver/mysql"
)

const dirverName = "mysql"

var log = logutil.NewPackageLogger()
