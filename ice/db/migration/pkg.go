// Package migration provides database schema migration
package migration // import "github.com/dyweb/go.ice/ice/db/migration"

import (
	"github.com/dyweb/go.ice/ice/util/logutil"
)

const initTaskName = "create_migration_table"
const migrationTableName = "icemigration"

var log = logutil.NewPackageLogger()
