package migration

import (
	"github.com/at15/go.ice/ice/util/logutil"
)

// InitialMigration is an reserved migration task, which creates migration table
const InitialMigration = "_initial"

var log = logutil.NewPackageLogger()
