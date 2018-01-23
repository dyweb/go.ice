package cmd

import (
	"database/sql"
	"github.com/at15/go.ice/ice/db/migration"
	"github.com/spf13/cobra"
)

func makeMigrationCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Long:  "Run registered migration tasks to update schema and feed fixture",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.PreRun(dbc, cmd, args)
			var (
				tx  *sql.Tx
				err error
			)
			w := dbc.MustWrapper()
			if tx, err = w.Transaction(); err != nil {
				log.Fatal(err)
			}
			init := migration.InitTask()
			if err = init.Up(tx); err != nil {
				log.Fatal(err)
			} else {
				if err = tx.Commit(); err != nil {
					log.Fatalf("failed to commit %v", err)
				}
			}
			log.Info("migration finished")
			// TODO: dbc should have cleanup, close connection etc.
			// Aborted connection 6 to db: 'icehub' user: 'root' host: '172.19.0.1' (Got an error reading communication packets)
		},
	}
}
