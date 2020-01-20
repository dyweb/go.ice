package cli

import (
	"io/ioutil"
	"os"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// root.go is the root command

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
}

func (r *Root) Command() *cobra.Command {
	if r.cmd != nil {
		return r.cmd
	}
	r.makeRootCmd()
	r.cmd.AddCommand(makeVersionCmd(r))
	return r.cmd
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

func IsServer() func(app *Root) {
	return func(app *Root) {
		app.server = true
	}
}

func (r *Root) Name() string {
	return r.name
}

func (r *Root) Description() string {
	return r.description
}

func (r *Root) Version() string {
	return r.buildInfo.Version
}

func (r *Root) ConfigFile() string {
	return r.configFile
}

func (r *Root) LoadConfigTo(v interface{}) error {
	b, err := ioutil.ReadFile(r.configFile)
	if err != nil {
		return err
	}
	if err := yaml.UnmarshalStrict(b, v); err != nil {
		return errors.Wrap(err, "error decode config as YAML")
	}
	return nil
}

func (r *Root) makeRootCmd() {
	r.cmd = &cobra.Command{
		Use:   r.Name(),
		Short: r.Description(),
		Long:  r.Description(),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			// we exit 1 because user may pass nothing and hope it run, which is never the case for go.ice based app,
			// where the real logic is always in sub commands
			os.Exit(1)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Use == "version" || cmd.Use == r.Name() {
				return
			}
			if err := r.updateLogSettings(); err != nil {
				log.Fatal(err)
				return
			}
		},
	}
	r.bindFlags()
}

func (r *Root) updateLogSettings() error {
	// default log handler has no color
	if r.logColor {
		var h dlog.Handler
		if r.server {
			// server runs a long time, delta would overflow ...
			h = cli.New(os.Stderr, false)
		} else {
			// client cli normally prefer delta to show time elapsed
			h = cli.New(os.Stderr, true)
		}
		dlog.SetHandler(h)
	}
	if r.logSource {
		dlog.EnableSource()
	}
	if r.verbose {
		dlog.SetLevel(dlog.DebugLevel)
		// FIXME: registry might need to allow caller to have a first logger
		//r.logRegistry.Debug("using debug level logging due to verbose config")
	}
	return nil
}

func (r *Root) bindFlags() {
	cmd := r.cmd
	// config
	cmd.PersistentFlags().StringVar(&r.configFile, "config", r.Name()+".yml", "config file location")
	// log
	cmd.PersistentFlags().BoolVar(&r.verbose, "verbose", false, "verbose output and set log level to debug")
	cmd.PersistentFlags().BoolVar(&r.logSource, "logSrc", false, "log source line when logging (expensive)")
	cmd.PersistentFlags().BoolVar(&r.logColor, "logColor", true, "disable log color if set to false")
}
