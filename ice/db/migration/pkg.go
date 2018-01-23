package migration

import (
	"github.com/at15/go.ice/ice/util/logutil"
)

const initTaskName = "create_migration_table"
const migrationTableNameQuoted = "`_ice_migration`"

var log = logutil.NewPackageLogger()
