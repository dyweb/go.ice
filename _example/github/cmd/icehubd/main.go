package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// specified in makefile
var version string

var rootCmd = &cobra.Command{
	Use:   "icehubd",
	Short: "icehub daemon",
	Long:  "IceHub is an example GitHub integration service using go.ice",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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

// TODO: command for migrating database (create table, fill in dummy data)
// TODO: icehub ice, can cobar command be nested and have flag proper parsed?
func main() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
