package db

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// cobra command for database related operations

// Command is a wrapper to allow user update manager after cobra commands have been created
type Command struct {
	Root   *cobra.Command
	PreRun func(dbc *Command, cmd *cobra.Command, args []string)
	Mgr    *Manager
}

var rootCmd = &cobra.Command{
	Use:   "db",
	Short: "database maintenance",
	Long:  "Database drivers, migration, status, REPL",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

// TODO: dbs registered and dbs configured
var driverCmd = &cobra.Command{
	Use:   "drivers",
	Short: "registered database drivers",
	Long:  "Show registered database drivers",
	Run: func(cmd *cobra.Command, args []string) {
		drivers := Drivers()
		if len(drivers) == 0 {
			fmt.Println("not database/sql driver registered")
			return
		}
		fmt.Printf("%d drivers registered \n", len(drivers))
		for _, d := range drivers {
			fmt.Println(d)
		}
	},
}

func makeConfigCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "print configuration",
		Long:  "Print configuration of manager and databases",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.PreRun(dbc, cmd, args)
			dbc.Mgr.PrintConfig()
		},
	}
}

func makePingCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "check database connectivity",
		Long:  "Check if database is reachable",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.PreRun(dbc, cmd, args)
			var (
				w        *Wrapper
				duration time.Duration
				err      error
			)
			if len(args) > 0 {
				w, err = dbc.Mgr.Wrapper(args[0])
			} else {
				w, err = dbc.Mgr.Default()
			}
			if err != nil {
				log.Fatal(err)
			}
			if duration, err = w.Ping(5 * time.Second); err != nil {
				log.Fatal(err)
			} else {
				log.Infof("ping took %s", duration)
			}
		},
	}
}

// TODO: command for migrating database (create table, fill in dummy data)
// TODO: dbshell https://docs.djangoproject.com/en/2.0/ref/django-admin/#dbshell
// - also consider support docker container ...
func NewCommand(preRun func(dbc *Command, cmd *cobra.Command, args []string)) *Command {
	dbc := &Command{Mgr: nil, PreRun: preRun}
	root := *rootCmd
	root.AddCommand(driverCmd)
	root.AddCommand(makeConfigCmd(dbc))
	root.AddCommand(makePingCmd(dbc))
	dbc.Root = &root
	return dbc
}
