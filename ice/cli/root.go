package cli

import (
	"os"

	"github.com/dyweb/gommon/config"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
	"github.com/spf13/cobra"
)

type Root struct {
	cmd          *cobra.Command
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
	if root.cmd != nil {
		return root.cmd
	}
	root.cmd = makeRootCmd(root)
	root.cmd.AddCommand(makeVersionCmd(root))
	return root.cmd
}

func makeRootCmd(root *Root) *cobra.Command {
	cmd := &cobra.Command{
		Use:   root.Name(),
		Short: root.Description(),
		Long:  root.Description(),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			// we exit 1 because user may pass nothing and hope it run, which is never the case for go.ice based app,
			// where the real logic is always in sub commands
			os.Exit(1)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Use == "version" || cmd.Use == root.Name() {
				return
			}
			// TODO: user may forgot to set logRegistry in option, and this will cause panic on nil pointer
			dlog.SetHandlerRecursive(root.logRegistry, cli.New(os.Stderr, true))
			if root.logSource {
				dlog.EnableSourceRecusrive(root.logRegistry)
			}
			if root.verbose {
				dlog.SetLevelRecursive(root.logRegistry, dlog.DebugLevel)
				root.logRegistry.Debug("using debug level logging due to verbose config")
			}
		},
	}
	cmd.PersistentFlags().StringVar(&root.configFile, "config", root.Name()+".yml", "config file location")
	cmd.PersistentFlags().BoolVar(&root.verbose, "verbose", false, "verbose output and set log level to debug")
	cmd.PersistentFlags().BoolVar(&root.logSource, "logsrc", false, "log source line when logging (expensive)")
	return cmd
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
		return errors.Wrap(err, "can't load config file")
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
