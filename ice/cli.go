package ice

import (
	"fmt"
	"os"

	"github.com/dyweb/gommon/config"
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
)

type Cli struct {
	root         *cobra.Command
	name         string
	description  string
	buildInfo    BuildInfo
	config       interface{}
	configFile   string
	configLoaded bool
	verbose      bool
	logSource    bool
	logRegistry  *dlog.Logger
}

type BuildInfo struct {
	Version   string
	Commit    string
	BuildTime string
	BuildUser string
	GoVersion string
}

// use functional options https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

type CliOptions func(a *Cli)

func New(options ...CliOptions) *Cli {
	a := &Cli{
		config: nil,
	}
	for _, opt := range options {
		opt(a)
	}
	return a
}

func NewCmd(app *Cli) *cobra.Command {
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
			dlog.SetHandlerRecursive(app.logRegistry, cli.New(os.Stderr, true))
			if app.logSource {
				dlog.EnableSourceRecusrive(app.logRegistry)
			}
			if app.verbose {
				dlog.SetLevelRecursive(app.logRegistry, dlog.DebugLevel)
				app.logRegistry.Debug("using debug level logging due to verbose config")
			}
		},
	}
	root.PersistentFlags().StringVar(&app.configFile, "config", app.Name()+".yml", "config file location")
	root.PersistentFlags().BoolVar(&app.verbose, "verbose", false, "verbose output and set log level to debug")
	root.PersistentFlags().BoolVar(&app.logSource, "logsrc", false, "log source line when logging (expensive)")
	ver := &cobra.Command{
		Use:   "version",
		Short: "print version",
		Long:  "Print current version " + app.Version(),
		Run: func(cmd *cobra.Command, args []string) {
			if app.verbose {
				app.buildInfo.PrintTo(os.Stdout)
			} else {
				fmt.Println(app.Version())
			}
		},
	}
	root.AddCommand(ver)
	app.root = root
	return root
}

func Name(name string) func(app *Cli) {
	return func(app *Cli) {
		app.name = name
	}
}
func Description(desc string) func(app *Cli) {
	return func(app *Cli) {
		app.description = desc
	}
}

func Version(info BuildInfo) func(app *Cli) {
	return func(app *Cli) {
		app.buildInfo = info
	}
}

func LogRegistry(logger *dlog.Logger) func(app *Cli) {
	return func(app *Cli) {
		app.logRegistry = logger
	}
}

func (b *Cli) Name() string {
	return b.name
}

func (b *Cli) Description() string {
	return b.description
}

func (b *Cli) Version() string {
	return b.buildInfo.Version
}

func (b *Cli) Config() interface{} {
	if b.config == nil {
		b.logRegistry.Warn("application config is nil")
	}
	return b.config
}

// TODO: go.ice should handle loading the yaml, marshal etc. as well
func (b *Cli) ConfigFile() string {
	return b.configFile
}

// TODO: check config file using gommon config
// TODO: have a config reader struct instead of using static package level method
// TODO: config file also specify logging (which package to log etc.)
func (b *Cli) LoadConfigTo(cfg interface{}) error {
	if err := config.LoadYAMLAsStruct(b.configFile, cfg); err != nil {
		return errors.WithMessage(err, "can't load config file")
	}
	b.config = cfg
	b.configLoaded = true
	return nil
}

func (b *Cli) IsConfigLoaded() bool {
	return b.configLoaded
}

func (b *Cli) SetConfigLoaded() {
	b.configLoaded = true
}

func (info *BuildInfo) PrintTo(w io.Writer) {
	fmt.Fprintf(w, "version: %s\n", info.Version)
	fmt.Fprintf(w, "commit: %s\n", info.Commit)
	fmt.Fprintf(w, "build time: %s\n", info.BuildTime)
	fmt.Fprintf(w, "build user: %s\n", info.BuildUser)
	fmt.Fprintf(w, "go version: %s\n", info.GoVersion)

}
