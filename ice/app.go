package ice

import (
	"fmt"
	"os"

	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
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
	logRegistry  *dlog.Logger
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
			// we exit 1 because user may pass nothing and hope it run, which is never the case for go.ice based app
			// the real logic is always in sub commands
			os.Exit(1)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Use == "version" || cmd.Use == app.Name() {
				return
			}
			if app.verbose {
				dlog.SetLevelRecursive(app.logRegistry, dlog.DebugLevel)
				app.logRegistry.Debug("using debug level logging due to verbose config")
			}
		},
	}
	root.PersistentFlags().StringVar(&app.configFile, "config", app.Name()+".yml", "config file location")
	root.PersistentFlags().BoolVar(&app.verbose, "verbose", false, "verbose output and set log level to debug")
	ver := &cobra.Command{
		Use:   "version",
		Short: "print version",
		Long:  "Print current version " + app.Version(),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app.Version())
			if app.verbose {
				// TODO: print build info in verbose mode
				//			fmt.Printf("version: %s\n", common.Version())
				//			fmt.Printf("commit: %s\n", common.GitCommit())
				//			fmt.Printf("build time: %s\n", common.BuildTime())
				//			fmt.Printf("build user: %s\n", common.BuildUser())
			}
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

func LogRegistry(logger *dlog.Logger) func(app *App) {
	return func(app *App) {
		app.logRegistry = logger
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

// TODO: go.ice should handle loading the yaml, marshal etc. as well
func (b *App) ConfigFile() string {
	return b.configFile
}

func (b *App) IsConfigLoaded() bool {
	return b.configLoaded
}

func (b *App) SetConfigLoaded() {
	b.configLoaded = true
}
