package cli

import (
	"os"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
	"github.com/spf13/cobra"
)

type Root struct {
	cmd *cobra.Command

	// set by flags
	verbose bool
	// TODO: might allow setting log level as well
	logSource  bool
	logColor   bool
	configFile string

	// set by compiler
	buildInfo BuildInfo

	// set by developer in main.go
	name        string
	description string
	server      bool
	logRegistry *dlog.Registry
}

func (root *Root) Command() *cobra.Command {
	if root.cmd != nil {
		return root.cmd
	}
	root.makeRootCmd()
	root.cmd.AddCommand(makeVersionCmd(root))
	// TODO: add gommon and database
	return root.cmd
}

// use functional options https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

type Options func(a *Root)

func New(options ...Options) *Root {
	root := Root{}
	for _, opt := range options {
		opt(&root)
	}
	return &root
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

func LogRegistry(logger *dlog.Registry) func(app *Root) {
	return func(app *Root) {
		app.logRegistry = logger
	}
}

func IsServer() func(app *Root) {
	return func(app *Root) {
		app.server = true
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

func (root *Root) ConfigFile() string {
	return root.configFile
}

func (root *Root) makeRootCmd() {
	root.cmd = &cobra.Command{
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
			if err := root.updateLogSettings(); err != nil {
				log.Fatal(err)
				return
			}
		},
	}
	root.bindFlags()
}

func (root *Root) updateLogSettings() error {
	if root.logRegistry == nil {
		// TODO: might show full solution for error, has a internal knowledge database
		return errors.New("logRegistry is not set for Root command, pass cli.LogRegistry(logger) when call cli.New")
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
		dlog.SetHandler(root.logRegistry, h)
	}
	if root.logSource {
		dlog.EnableSource(root.logRegistry)
	}
	if root.verbose {
		dlog.SetLevel(root.logRegistry, dlog.DebugLevel)
		// FIXME: registry might need to allow caller to have a first logger
		//root.logRegistry.Debug("using debug level logging due to verbose config")
	}
	return nil
}

func (root *Root) bindFlags() {
	cmd := root.cmd
	// config
	cmd.PersistentFlags().StringVar(&root.configFile, "config", root.Name()+".yml", "config file location")
	// log
	cmd.PersistentFlags().BoolVar(&root.verbose, "verbose", false, "verbose output and set log level to debug")
	cmd.PersistentFlags().BoolVar(&root.logSource, "logSrc", false, "log source line when logging (expensive)")
	cmd.PersistentFlags().BoolVar(&root.logColor, "logColor", true, "disable log color if set to false")
}
