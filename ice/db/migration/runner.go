package migration

import (
	"github.com/dyweb/gommon/errors"

	"github.com/at15/go.ice/ice/db"
)

type Direction bool

const (
	Up   Direction = true
	Down Direction = false
)

type TaskRunner struct {
	db *db.Wrapper
}

func NewRunner(db *db.Wrapper) *TaskRunner {
	return &TaskRunner{
		db: db,
	}
}

func (r *TaskRunner) Run(task Task, direction Direction) error {
	tx, err := r.db.Transaction()
	if err != nil {
		return errors.Wrap(err, "can't start transaction for migration")
	}
	if direction == Up {
		err = task.Up(tx, r.db.Adapter())
	} else {
		err = task.Down(tx, r.db.Adapter())
	}
	if err != nil {
		// it seems you may rollback ddl in some database
		// https://stackoverflow.com/questions/4692690/is-it-possible-to-roll-back-create-table-and-alter-table-statements-in-major-sql
		if rerr := tx.Rollback(); rerr != nil {
			return errors.Wrapf(rerr, "failed to rollback a failed migration %s %s %v", task.Name(), direction, err)
		} else {
			log.Infof("rollback failed migration %s %s %v", task.Name(), direction, err)
		}
		return errors.Wrapf(err, "failed to migrate %s %s", task.Name(), direction)
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "failed to commit migration %s %s", task.Name(), direction)
	}
	return nil
}

func (d Direction) String() string {
	if d == Up {
		return "up"
	} else {
		return "down"
	}
}
