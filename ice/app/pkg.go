// Package app define common interface for both client and server application
package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// TODO: we may not even need to distinguish client and server, that's too much wrapper

//type App interface {
//	// might be app config
//	//IsClient() bool
//	//IsServer() bool
//	//Verbose() bool
//	//Version() string
//}

// TODO: build info, as a struct?

type App interface {
	Name() string
	Description() string
	Version() string
}

func New(options ...BaseAppOptions) *BaseApp {
	a := &BaseApp{}
	for _, opt := range options {
		opt(a)
	}
	return a
}

// TODO: might just accept BaseApp struct directly ...
func NewCmd(app App) *cobra.Command {
	root := &cobra.Command{
		Use:   app.Name(),
		Short: app.Description(),
		Long:  app.Description(),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			// TODO: should we exit 1 or 0
			os.Exit(1)
		},
	}
	ver := &cobra.Command{
		Use:   "version",
		Short: "print version",
		Long:  "Print current version " + app.Version(),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app.Version())
			// TODO: verbose, use build info
		},
	}
	root.AddCommand(ver)
	return root
}

func Run(app App) {
	// TODO: common configuration?
	// TODO: handle sig term etc.
}
