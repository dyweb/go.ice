package main

import (
	"fmt"
	"os"

	"github.com/dyweb/gommon/config"
	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"

	"github.com/at15/go.ice/_example/github/pkg/server"
	"github.com/at15/go.ice/_example/github/pkg/util/logutil"
	"github.com/at15/go.ice/ice/db"

	_ "github.com/at15/go.ice/ice/db/adapters/sqlite"
	_ "github.com/at15/go.ice/ice/db/adapters/postgres"
	_ "github.com/at15/go.ice/ice/db/adapters/mysql"
)

//_ "github.com/go-sql-driver/mysql"

const (
	myname = "icehubd" // 你的名字
)

var log = logutil.Registry

// TODO: flags for enable debug logging etc. it should also be passed to sub commands like db

// specified in makefile
var version string

// specified using flags
var cfgFile string
var verbose = false

// global configuration instance
var cfgLoaded = false
var cfg server.Config

// TODO: might need a registry of application instead of scatter variables around in main
var dbMgr *db.Manager

var rootCmd = &cobra.Command{
	Use:   myname,
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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Use == "version" || cmd.Use == myname {
			return
		}
		if verbose {
			dlog.SetLevelRecursive(log, dlog.DebugLevel)
			log.Debug("using debug level due to verbose flag")
		}
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

// TODO: just here for testing out the log command, though it might possible to make it like db command to be part
// of go.ice's built in command for managing common config
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "test log config",
	Long:  "Test log tree printer etc.",
	Run: func(cmd *cobra.Command, args []string) {
		log.PrintTree()
	},
}

// TODO: check config file using gommon config
func loadConfig() {
	if !cfgLoaded {
		// TODO: have a config reader struct instead of using static package level method
		// TODO: config file also specify logging (which package to log etc.)
		if err := config.LoadYAMLAsStruct(cfgFile, &cfg); err != nil {
			// TODO: use log
			fmt.Printf("can't load %s %v\n", cfgFile, err)
		}
		cfgLoaded = true
	}
}

// TODO: icehub ice, can cobra command be nested and have flag proper parsed? icehub db is there ...
func main() {
	// TODO: common root command should be put into a struct, but need another struct to store the flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "icehub.yml", "config file location")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output and set log level to debug")
	rootCmd.AddCommand(versionCmd)
	dbc := db.NewCommand(func(dbc *db.Command, cmd *cobra.Command, args []string) {
		if dbMgr == nil {
			loadConfig()
			dbMgr = db.NewManager(cfg.DatabaseManager)
		}
		if dbc.Mgr == nil {
			dbc.Mgr = dbMgr
		}
	})
	rootCmd.AddCommand(dbc.Root)
	rootCmd.AddCommand(logCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
