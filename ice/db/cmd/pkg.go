package cmd

import (
	"github.com/spf13/cobra"

	"github.com/at15/go.ice/ice/db"
	"github.com/at15/go.ice/ice/util/logutil"
)

var log = logutil.NewPackageLogger()

// Command is a wrapper to allow user update manager after cobra commands have been created
type Command struct {
	Root   *cobra.Command
	PreRun func(dbc *Command, cmd *cobra.Command, args []string) // TODO: change pre run? also need clean up ...
	Mgr    *db.Manager
	db     string // database selected by user
}

// TODO: command for migrating database (create table, fill in dummy data)
// TODO: dbshell https://docs.djangoproject.com/en/2.0/ref/django-admin/#dbshell
// - also consider support docker container ...
// TODO: create database (the user need to be root ... but this is normally the case in local dev ...)
func NewCommand(preRun func(dbc *Command, cmd *cobra.Command, args []string)) *Command {
	dbc := &Command{Mgr: nil, PreRun: preRun}
	root := *rootCmd
	// flags
	root.PersistentFlags().StringVar(&dbc.db, "db", "", "database to run command on, ping/migrate etc.")
	// sub commands
	root.AddCommand(driverCmd)
	root.AddCommand(adapterCmd)
	root.AddCommand(makeConfigCmd(dbc))
	root.AddCommand(makePingCmd(dbc))
	root.AddCommand(makeMigrationCmd(dbc))
	dbc.Root = &root
	return dbc
}

func (dbc *Command) MustWrapper() *db.Wrapper {
	var (
		w    *db.Wrapper
		name string
		err  error
	)
	if dbc.db != "" {
		name = dbc.db
	} else if name, err = dbc.Mgr.DefaultName(); err != nil {
		log.Fatal(err)
	}
	if w, err = dbc.Mgr.Wrapper(name); err != nil {
		log.Fatal(err)
	}
	return w
}
