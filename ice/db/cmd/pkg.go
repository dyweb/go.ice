package cmd

import (
	"github.com/dyweb/gommon/errors"
	"github.com/spf13/cobra"

	"github.com/dyweb/go.ice/ice/config"
	"github.com/dyweb/go.ice/ice/db"
	"github.com/dyweb/go.ice/ice/util/logutil"
)

// TODO: command for migrating database (create table, fill in dummy data)
// TODO: dbshell https://docs.djangoproject.com/en/2.0/ref/django-admin/#dbshell, also consider support docker container
// TODO: create/drop database (the user need to be root ... but this is normally the case in local dev ...)
// TODO: util function for clean up manager

var log, _ = logutil.NewPackageLoggerAndRegistry()

// Command is a wrapper to keep internal states like database manager
type Command struct {
	db           string // database selected by user
	configLoader func() (config.DatabaseManagerConfig, error)
	manager      *db.Manager
	root         *cobra.Command
}

func New(configLoader func() (config.DatabaseManagerConfig, error)) *Command {
	dbc := &Command{
		manager:      nil,
		configLoader: configLoader,
	}
	root := *rootCmd
	// flags
	root.PersistentFlags().StringVar(&dbc.db, "db", "", "database to run command on, ping/migrate etc.")
	// sub commands
	root.AddCommand(driverCmd, adapterCmd,
		makeConfigCmd(dbc), makePingCmd(dbc), makeMigrationCmd(dbc), makeCreateCmd(dbc))
	dbc.root = &root
	return dbc
}

func (dbc *Command) Root() *cobra.Command {
	return dbc.root
}

func (dbc *Command) mustConfigManager() {
	if err := dbc.configManager(); err != nil {
		log.Fatal(err)
	}
}

func (dbc *Command) configManager() error {
	if dbc.manager != nil {
		log.Debug("manager is already configured")
		return nil
	}
	if c, err := dbc.configLoader(); err != nil {
		return errors.Wrap(err, "can't load config to create manager")
	} else {
		dbc.manager = db.NewManager(c)
		return nil
	}
}

func (dbc *Command) mustWrapper(useDatabase bool) *db.Wrapper {
	var (
		w    *db.Wrapper
		name string
		err  error
	)
	if dbc.db != "" {
		name = dbc.db
	} else if name, err = dbc.manager.DefaultName(); err != nil {
		log.Fatal(err)
	}
	if w, err = dbc.manager.Wrapper(name, useDatabase); err != nil {
		log.Fatal(err)
	}
	return w
}

func (dbc *Command) close() {
	if dbc.manager == nil {
		return
	}
	if err := dbc.manager.Close(); err != nil {
		log.Warnf("error when closing database %v", err)
	} else {
		log.Info("database closed")
	}
}
