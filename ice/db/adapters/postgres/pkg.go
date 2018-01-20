package postgres

import (
	"github.com/at15/go.ice/ice/db"
	"github.com/at15/go.ice/ice/util/logutil"

	_ "github.com/jackc/pgx/stdlib" // TODO: pgx also support its native access, and how is JSONB handled?
)

const adapterName = "postgres"
const driverName = "pgx"

var log = logutil.NewPackageLogger()

func init() {
	db.RegisterAdapterFactory(adapterName, func() db.Adapter {
		return New()
	})
}
