package main

import (
	"fmt"
	"os"

	"github.com/dyweb/gommon/config"
	"github.com/spf13/cobra"

	"github.com/at15/go.ice/_example/github/pkg/server"
	"github.com/at15/go.ice/ice/db"
	_ "github.com/jackc/pgx/stdlib" // TODO: pgx also support its native access, and how is JSONB handled
	_ "github.com/mattn/go-sqlite3" // nameless import to register driver
)

//_ "github.com/go-sql-driver/mysql"

// TODO: flags for enable debug logging etc. it should also be passed to sub commands like db
// TODO: load and check config file using gommon config

// specified in makefile
var version string

// specified using flags
var cfgFile string

// global configuration instance
var cfgLoaded = false
var cfg server.Config
var dbMgr *db.Manager

var rootCmd = &cobra.Command{
	Use:   "icehubd",
	Short: "icehub daemon",
	Long:  "IceHub is an example GitHub integration service using go.ice",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
		//IceHub is an example GitHub integration service using go.ice
		//
		//Usage:
		//	icehubd [flags]
		//	icehubd [command]
		//
		//	Available Commands:
		//	help        Help about any command
		//	version     print version
		//
		//Flags:
		//	-h, --help   help for icehubd
		//
		//Use "icehubd [command] --help" for more information about a command.

		// usage does not have the long description like help
		//cmd.Usage()
		//Usage:
		//	icehubd [flags]
		//	icehubd [command]
		//
		//	Available Commands:
		//	help        Help about any command
		//	version     print version
		//
		//Flags:
		//	-h, --help   help for icehubd
		//
		//	Use "icehubd [command] --help" for more information about a command.
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  "Print current version " + version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func loadConfig() {
	if !cfgLoaded {
		// TODO: set logging level based on flag
		// TODO: config file also specify logging (which package to log etc.)
		if err := config.LoadYAMLAsStruct(cfgFile, &cfg); err != nil {
			// TODO: use log
			fmt.Printf("can't load %s %v\n", cfgFile, err)
		}
		cfgLoaded = true
	}
}

// TODO: icehub ice, can cobra command be nested and have flag proper parsed?
func main() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "icehub.yml", "config file location")
	rootCmd.AddCommand(versionCmd)
	dbc := db.NewCommand(func(dbc *db.Command, cmd *cobra.Command, args []string) {
		if dbc.Mgr == nil {
			loadConfig()
			dbc.Mgr = db.NewManager(cfg.DatabaseManager)
		}
	})
	rootCmd.AddCommand(dbc.Root)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
