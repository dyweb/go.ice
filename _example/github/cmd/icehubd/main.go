package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/at15/go.ice/ice/db"

	_ "github.com/mattn/go-sqlite3" // nameless import to register driver
	_ "github.com/jackc/pgx/stdlib" // TODO: pgx also support its native access, and how is JSONB handled
	_ "github.com/go-sql-driver/mysql"
)

// TODO: flags for enable debug logging etc. it should also be passed to sub commands like db

// specified in makefile
var version string

var rootCmd = &cobra.Command{
	Use:   "icehubd",
	Short: "icehub daemon",
	Long:  "IceHub is an example GitHub integration service using go.ice",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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

// TODO: icehub ice, can cobar command be nested and have flag proper parsed?
func main() {
	rootCmd.AddCommand(versionCmd)
	// TODO: initialize db manager based on config
	dbCmd := db.NewCommand(nil)
	rootCmd.AddCommand(dbCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
