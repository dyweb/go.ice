package ice

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

// TODO: build info, as a struct?
type App struct {
	root         *cobra.Command
	name         string
	description  string
	version      string
	configFile   string
	configLoaded bool
	verbose      bool
}

// use functional options https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

type AppOptions func(a *App)

func New(options ...AppOptions) *App {
	a := &App{}
	for _, opt := range options {
		opt(a)
	}
	return a
}

func NewCmd(app *App) *cobra.Command {
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
	app.root = root
	return root
}

func Name(name string) func(app *App) {
	return func(app *App) {
		app.name = name
	}
}
func Description(desc string) func(app *App) {
	return func(app *App) {
		app.description = desc
	}
}

func Version(ver string) func(app *App) {
	return func(app *App) {
		app.version = ver
	}
}

func (b *App) Name() string {
	return b.name
}

func (b *App) Description() string {
	return b.description
}

func (b *App) Version() string {
	return b.version
}
