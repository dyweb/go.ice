package sqlite

import (
	"github.com/at15/go.ice/ice/util/logutil"

	_ "github.com/mattn/go-sqlite3" // nameless import to register driver
)

var log = logutil.NewPackageLogger()
