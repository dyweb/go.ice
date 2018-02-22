package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func makePingCmd(dbc *Command) *cobra.Command {
	useDatabase := true
	c := &cobra.Command{
		Use:   "ping",
		Short: "check database connectivity",
		Long:  "Check if database is reachable",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.mustConfigManager()
			w := dbc.mustWrapper(useDatabase)
			if duration, err := w.Ping(5 * time.Second); err != nil {
				log.Fatal(err)
			} else {
				log.Infof("ping took %s", duration)
			}
			dbc.close()
		},
	}
	//icehubd db ping --usedb=false
	//INFO 0000 ping took 1.856071ms
	//INFO 0000 database closed
	//icehubd db ping --usedb=true
	//FATA 0000 Error 1049: Unknown database 'icehub'
	c.Flags().BoolVar(&useDatabase, "usedb", true, "use database name when connect")
	return c
}
