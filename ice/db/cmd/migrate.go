package cmd

import (
	"github.com/at15/go.ice/ice/db/migration"
	"github.com/spf13/cobra"
)

func makeMigrationCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Long:  "Run registered migration tasks to update schema and feed fixture",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.mustConfigManager()
			w := dbc.mustWrapper(true)
			runner := migration.NewRunner(w)
			log.Info("runner created!")
			// TODO: check if migration table exists
			if err := runner.Run(migration.InitTask(), migration.Up); err != nil {
				dbc.close()
				log.Fatal(err)
			}
			log.Info("migration finished")
			// Aborted connection 6 to db: 'icehub' user: 'root' host: '172.19.0.1' (Got an error reading communication packets)
			dbc.close()
		},
	}
}
