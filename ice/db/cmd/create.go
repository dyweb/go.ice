package cmd

import "github.com/spf13/cobra"

func makeCreateCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create database",
		Long:  "Create database",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.mustConfigManager()
			w := dbc.mustWrapper(false)
			cfg, err := dbc.manager.DefaultConfig()
			if err != nil {
				dbc.close()
				log.Fatal(err)
			}
			if err := w.CreateDatabase(cfg.DBName); err != nil {
				dbc.close()
				log.Fatal(err)
			}
			log.Infof("database %s created", cfg.DBName)
			dbc.close()
		},
	}
}
