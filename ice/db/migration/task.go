package migration

import (
	"database/sql"
	"time"

	"fmt"
	"github.com/pkg/errors"
)

var initialTask = NewTask(
	time.Date(2018, 1, 21, 23, 43, 01, 0, time.UTC),
	initTaskName, "create migration table to track future migration tasks",
	createMigrationTable, dropMigrationTable)

// blankInitialTask is used to break initialization loop
var blankInitialTask = NewTask(
	time.Date(2018, 1, 21, 23, 43, 01, 0, time.UTC),
	initTaskName, "create migration table to track future migration tasks",
	nil, nil)

type TaskFunc func(tx *sql.Tx) error

type Task interface {
	CreateTime() time.Time
	Name() string
	Description() string
	// Up does not need to commit or rollback, runner with handle it based on error
	Up(tx *sql.Tx) error
	// Down does not need to commit or rollback, runner with handle it based on error
	Down(tx *sql.Tx) error
}

type BasicTask struct {
	createTime  time.Time
	name        string
	description string
	up          TaskFunc
	down        TaskFunc
}

func InitTask() Task {
	return initialTask
}

func NewTask(t time.Time, name string, desc string, up TaskFunc, down TaskFunc) Task {
	return &BasicTask{
		createTime:  t,
		name:        name,
		description: desc,
		up:          up,
		down:        down,
	}
}

func (t *BasicTask) CreateTime() time.Time {
	return t.createTime
}

func (t *BasicTask) Name() string {
	return t.name
}

func (t *BasicTask) Description() string {
	return t.description
}

func (t *BasicTask) Up(tx *sql.Tx) error {
	return t.up(tx)
}

func (t *BasicTask) Down(tx *sql.Tx) error {
	return t.down(tx)
}

func createMigrationTable(tx *sql.Tx) error {
	// we need to use ` to quote the table name `_ice_migration`, which is why we concat string instead of using literal
	c := "CREATE TABLE " + migrationTableName + " (" +
		` name VARCHAR(255), description TEXT,
		  create_time INT, apply_time INT, update_time INT,
          status INT);`
	if _, err := tx.Exec(c); err != nil {
		return errors.Wrap(err, "can't create migration table")
	}
	// FIXME: the task info is not inserted? the table is created though ...
	return insertTaskInfo(tx, blankInitialTask)
}

func insertTaskInfo(tx *sql.Tx, task Task) error {
	log.Info("insert task info!")
	now := time.Now().Unix()
	// TODO: prepare statement syntax varies based on database
	stmt, err := tx.Prepare(fmt.Sprintf(
		"INSERT INTO %s (name, description, create_time, apply_time, update_time, status) VALUES (?, ?, ?, ?, ?, ?)", migrationTableName))
	if err != nil {
		return errors.Wrap(err, "can't prepare statement")
	}
	defer stmt.Close()
	if _, err := stmt.Exec(task.Name(), task.Description(), task.CreateTime().Unix(), now, now, Success); err != nil {
		return errors.Wrap(err, "can't insert migration record")
	}
	return nil
}

func dropMigrationTable(tx *sql.Tx) error {
	d := "DROP TABLE " + migrationTableName
	if _, err := tx.Exec(d); err != nil {
		return errors.Wrap(err, "can't drop migration table")
	}
	return nil
}
