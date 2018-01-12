package db

import (
	"github.com/spf13/cobra"
)

// cobra command for database related operations

func makeRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "db",
		Short: "database maintenance",
		Long:  "Database drivers, migration, status, REPL",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}

// TODO: command for migrating database (create table, fill in dummy data)

func NewCommand(mgr *Manager) *cobra.Command {
	root := makeRootCommand()
	return root
}
