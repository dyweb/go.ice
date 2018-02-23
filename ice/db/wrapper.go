package db

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/at15/go.ice/ice/config"
	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"
)

type Wrapper struct {
	db  *sql.DB
	a   Adapter
	c   config.DatabaseConfig
	mu  sync.RWMutex
	log *dlog.Logger
}

func NewWrapper(a Adapter) *Wrapper {
	w := &Wrapper{a: a}
	w.log = dlog.NewStructLogger(log, w)
	return w
}

func (w *Wrapper) Adapter() Adapter {
	return w.a
}

func (w *Wrapper) SetDB(db *sql.DB) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if db == nil {
		w.log.Warn("sql.DB pointer is nil")
	}
	w.db = db
}

// TODO: add error in return value? now we just warn if its nil ...
func (w *Wrapper) GetDB() *sql.DB {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.db == nil {
		w.log.Warn("sql.DB pointer is nil")
	}
	return w.db
}

func (w *Wrapper) Transaction() (*sql.Tx, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.db == nil {
		return nil, errors.New("sql.DB pointer is nil")
	}
	if tx, err := w.db.BeginTx(context.Background(), nil); err != nil {
		return tx, errors.Wrap(err, "can't begin transaction")
	} else {
		return tx, nil
	}
}

func (w *Wrapper) Ping(timeout time.Duration) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	start := time.Now()
	err := w.db.PingContext(ctx)
	duration := time.Now().Sub(start)
	if err != nil {
		return duration, errors.WithStack(err)
	}
	return duration, nil
}

func (w *Wrapper) CreateDatabase(name string) error {
	if !w.a.CanCreateDatabase() {
		log.Warnf("%s does not support create database", w.a.DriverName())
		return nil
	}
	db := w.db
	// NOTE: you can't use transaction for create database ...
	// pg: ERROR:  CREATE DATABASE cannot run inside a transaction block
	if _, err := db.Exec("CREATE DATABASE " + name + ";"); err != nil {
		return errors.Wrapf(err, "can't create database %s", name)
	}
	log.Debugf("database %s created", name)
	return nil
}

func (w *Wrapper) Close() error {
	if w.db == nil {
		w.log.Warn("closing nil db")
		return nil
	}
	if err := w.db.Close(); err != nil {
		return errors.Wrap(err, "error when closing db")
	}
	return nil
}
