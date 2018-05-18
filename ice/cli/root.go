package cli

import (
	"os"

	"github.com/dyweb/gommon/config"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
	"github.com/spf13/cobra"
)

type root struct {
	cmd *cobra.Command

	config       interface{}
	configFile   string
	configLoaded bool

	// set by flags
	verbose   bool
	logSource bool
	logColor  bool

	// set by compiler
	buildInfo BuildInfo

	// set by developer in main.go
	name        string
	description string
	server      bool
	logRegistry *dlog.Logger
}

// use functional options https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

type Options func(a *root)

func New(options ...Options) *root {
	root := &root{
		config: nil,
	}
	for _, opt := range options {
		opt(root)
	}
	return root
}

func (root *root) Command() *cobra.Command {
	if root.cmd != nil {
		return root.cmd
	}
	root.cmd = makeRootCmd(root)
	root.cmd.AddCommand(makeVersionCmd(root))
	// TODO: add gommon and database
	return root.cmd
}

func makeRootCmd(root *root) *cobra.Command {
	cmd := cobra.Command{
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
			if root.logRegistry == nil {
				log.Fatal("logRegistry is not set for root command, pass cli.LogRegistry(logger) when call cli.New")
				return
			}
			// default log handler has no color
			if root.logColor {
				var h dlog.Handler
				if root.server {
					// server runs a long time, delta would overflow ...
					h = cli.New(os.Stderr, false)
				} else {
					// client cli normally prefer delta to show time elapsed
					h = cli.New(os.Stderr, true)
				}
				dlog.SetHandlerRecursive(root.logRegistry, h)
			}
			if root.logSource {
				dlog.EnableSourceRecursive(root.logRegistry)
			}
			if root.verbose {
				dlog.SetLevelRecursive(root.logRegistry, dlog.DebugLevel)
				root.logRegistry.Debug("using debug level logging due to verbose config")
			}
		},
	}
	cmd.PersistentFlags().StringVar(&root.configFile, "config", root.Name()+".yml", "config file location")
	cmd.PersistentFlags().BoolVar(&root.verbose, "verbose", false, "verbose output and set log level to debug")
	cmd.PersistentFlags().BoolVar(&root.logSource, "logSrc", false, "log source line when logging (expensive)")
	cmd.PersistentFlags().BoolVar(&root.logColor, "logColor", true, "disable log color if set to false")
	return &cmd
}

func Name(name string) func(app *root) {
	return func(app *root) {
		app.name = name
	}
}
func Description(desc string) func(app *root) {
	return func(app *root) {
		app.description = desc
	}
}

func Version(info BuildInfo) func(app *root) {
	return func(app *root) {
		app.buildInfo = info
	}
}

func LogRegistry(logger *dlog.Logger) func(app *root) {
	return func(app *root) {
		app.logRegistry = logger
	}
}

func IsServer() func(app *root) {
	return func(app *root) {
		app.server = true
	}
}

func (root *root) Name() string {
	return root.name
}

func (root *root) Description() string {
	return root.description
}

func (root *root) Version() string {
	return root.buildInfo.Version
}

func (root *root) Config() interface{} {
	if root.config == nil {
		root.logRegistry.Warn("application config is nil")
	}
	return root.config
}

func (root *root) ConfigFile() string {
	return root.configFile
}

// TODO: have a config reader struct instead of using static package level method
// TODO: config file also specify logging (which package to log etc.)
func (root *root) LoadConfigTo(cfg interface{}) error {
	if err := config.LoadYAMLAsStruct(root.configFile, cfg); err != nil {
		return errors.Wrap(err, "can't load config file")
	}
	return root.loadConfig(cfg)
}

func (root *root) LoadConfigToStrict(cfg interface{}) error {
	if err := config.LoadYAMLDirectStrict(root.configFile, cfg); err != nil {
		return errors.Wrap(err, "can't load config file in strict mode, check typos")
	}
	return root.loadConfig(cfg)
}

func (root *root) loadConfig(cfg interface{}) error {
	root.config = cfg
	root.configLoaded = true
	return nil
}

func (root *root) IsConfigLoaded() bool {
	return root.configLoaded
}

func (root *root) SetConfigLoaded() {
	root.configLoaded = true
}
