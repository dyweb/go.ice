package db

import (
	"github.com/spf13/cobra"
	"fmt"
)

// cobra command for database related operations

var rootCmd = &cobra.Command{
	Use:   "db",
	Short: "database maintenance",
	Long:  "Database drivers, migration, status, REPL",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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

func makeConfigCmd(mgr *Manager) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "print configuration",
		Long:  "Print configuration of manager and databases",
		Run: func(cmd *cobra.Command, args []string) {
			mgr.PrintConfig()
		},
	}
}

// TODO: command for migrating database (create table, fill in dummy data)

func NewCommand(mgr *Manager) *cobra.Command {
	root := *rootCmd
	root.AddCommand(driverCmd)
	root.AddCommand(makeConfigCmd(mgr))
	return &root
}
