package postgres

import (
	"github.com/at15/go.ice/ice/util/logutil"

	_ "github.com/jackc/pgx/stdlib" // TODO: pgx also support its native access, and how is JSONB handled?
)

var log = logutil.NewPackageLogger()
