package db

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

var driverCmd = &cobra.Command{
	Use:   "drivers",
	Short: "registered database drivers",
	Long:  "Show registered database drivers",
	Run: func(cmd *cobra.Command, args []string) {
		drivers := Drivers()
		if len(drivers) == 0 {
			fmt.Println("not driver registered for database/sql")
			return
		}
		fmt.Printf("%d drivers found\n", len(drivers))
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

// TODO: command for migrating database (create table, fill in dummy data)
// TODO: dbshell https://docs.djangoproject.com/en/2.0/ref/django-admin/#dbshell
// - also consider support docker container ...
func NewCommand(preRun func(dbc *Command, cmd *cobra.Command, args []string)) *Command {
	dbc := &Command{Mgr: nil, PreRun: preRun}
	root := *rootCmd
	root.AddCommand(driverCmd)
	root.AddCommand(makeConfigCmd(dbc))
	dbc.Root = &root
	return dbc
}
