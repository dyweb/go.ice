package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/dyweb/gommon/config"
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
)

type Root struct {
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

type Options func(a *Root)

func New(options ...Options) *Root {
	root := &Root{
		config: nil,
	}
	for _, opt := range options {
		opt(root)
	}
	return root
}

func (root *Root) Command() *cobra.Command {
	return root.root
}

func NewCmd(app *Root) *cobra.Command {
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

func Name(name string) func(app *Root) {
	return func(app *Root) {
		app.name = name
	}
}
func Description(desc string) func(app *Root) {
	return func(app *Root) {
		app.description = desc
	}
}

func Version(info BuildInfo) func(app *Root) {
	return func(app *Root) {
		app.buildInfo = info
	}
}

func LogRegistry(logger *dlog.Logger) func(app *Root) {
	return func(app *Root) {
		app.logRegistry = logger
	}
}

func (root *Root) Name() string {
	return root.name
}

func (root *Root) Description() string {
	return root.description
}

func (root *Root) Version() string {
	return root.buildInfo.Version
}

func (root *Root) Config() interface{} {
	if root.config == nil {
		root.logRegistry.Warn("application config is nil")
	}
	return root.config
}

func (root *Root) ConfigFile() string {
	return root.configFile
}

// TODO: check config file using gommon config
// TODO: have a config reader struct instead of using static package level method
// TODO: config file also specify logging (which package to log etc.)
func (root *Root) LoadConfigTo(cfg interface{}) error {
	if err := config.LoadYAMLAsStruct(root.configFile, cfg); err != nil {
		return errors.WithMessage(err, "can't load config file")
	}
	root.config = cfg
	root.configLoaded = true
	return nil
}

func (root *Root) IsConfigLoaded() bool {
	return root.configLoaded
}

func (root *Root) SetConfigLoaded() {
	root.configLoaded = true
}

func (info *BuildInfo) PrintTo(w io.Writer) {
	fmt.Fprintf(w, "version: %s\n", info.Version)
	fmt.Fprintf(w, "commit: %s\n", info.Commit)
	fmt.Fprintf(w, "build time: %s\n", info.BuildTime)
	fmt.Fprintf(w, "build user: %s\n", info.BuildUser)
	fmt.Fprintf(w, "go version: %s\n", info.GoVersion)
}
