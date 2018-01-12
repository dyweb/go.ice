package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// specified in makefile
var version string

var rootCmd = &cobra.Command{
	Use:   "icehubctl",
	Short: "icehub client",
	Long:  "Client of IceHub, which is an example GitHub integration service using go.ice",
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

func main() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
