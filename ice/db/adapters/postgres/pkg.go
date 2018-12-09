// Package postgres wraps github.com/jackc/pgx
package postgres

import (
	"github.com/dyweb/go.ice/ice/db"
	"github.com/dyweb/go.ice/ice/util/logutil"

	_ "github.com/jackc/pgx/stdlib" // TODO: pgx also support its native access, and how is JSONB handled?
)

const adapterName = "postgres"
const driverName = "pgx"

var log, _ = logutil.NewPackageLoggerAndRegistry()

func init() {
	db.RegisterAdapterFactory(adapterName, func() db.Adapter {
		return New()
	})
}
