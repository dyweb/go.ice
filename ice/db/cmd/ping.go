package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func makePingCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "check database connectivity",
		Long:  "Check if database is reachable",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.mustConfigManager()
			w := dbc.mustWrapper()
			if duration, err := w.Ping(5 * time.Second); err != nil {
				log.Fatal(err)
			} else {
				log.Infof("ping took %s", duration)
			}
			// TODO: dbc should have cleanup
		},
	}
}
