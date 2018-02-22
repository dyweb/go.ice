package cmd

import "github.com/spf13/cobra"

func makeCreateCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create database",
		Long:  "Create database",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.mustConfigManager()
			dbc.mustWrapper(false)
			log.Info("TODO: create database")
			dbc.close()
		},
	}
}
