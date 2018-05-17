package migration

import (
	"fmt"
	"github.com/dyweb/go.ice/ice/db"
)

type Status int

const (
	Success  Status = 1
	Failed   Status = 2
	Rollback Status = 3
)

type Tracker struct {
	db *db.Wrapper
}

func NewTracker(db *db.Wrapper) *Tracker {
	return &Tracker{
		db: db,
	}
}

func (t *Tracker) GetExecuted() error {
	db := t.db.GetDB()
	db.Query(fmt.Sprintf("SELECT * FROM %s ORDER BY apply_time, create_time DESC WHERE status = %d", migrationTableName, Success))
	return nil
}
