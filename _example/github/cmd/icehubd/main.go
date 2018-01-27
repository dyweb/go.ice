package main

import (
	"fmt"
	"os"
	"time"
	"io"

	"github.com/opentracing/opentracing-go"
	jgconfig "github.com/uber/jaeger-client-go/config"
	"github.com/dyweb/gommon/config"
	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/at15/go.ice/_example/github/pkg/server"
	"github.com/at15/go.ice/_example/github/pkg/util/logutil"
	icfg "github.com/at15/go.ice/ice/config"
	"github.com/at15/go.ice/ice/db"
	idbcmd "github.com/at15/go.ice/ice/db/cmd"
	"github.com/at15/go.ice/_example/github/pkg/common"

	_ "github.com/at15/go.ice/ice/db/adapters/mysql"
	_ "github.com/at15/go.ice/ice/db/adapters/postgres"
	_ "github.com/at15/go.ice/ice/db/adapters/sqlite"
)

const (
	myname = "icehubd" // 你的名字
)

var log = logutil.Registry

// TODO: flags for enable debug logging etc. it should also be passed to sub commands like db

// specified using flags
var cfgFile string
var verbose = false

// global configuration instance
var cfgLoaded = false
var cfg server.Config

// TODO: might need a registry of application instead of scatter variables around in main
var dbMgr *db.Manager
var tracer opentracing.Tracer
var closer io.Closer

var rootCmd = &cobra.Command{
	Use:   myname,
	Short: "icehub daemon",
	Long:  "IceHub is an example GitHub integration service using go.ice",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Use == "version" || cmd.Use == myname {
			return
		}
		if verbose {
			dlog.SetLevelRecursive(log, dlog.DebugLevel)
			log.Debug("using debug level logging due to verbose flag")
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  "Print current version " + common.Version(),
	Run: func(cmd *cobra.Command, args []string) {
		if !verbose {
			fmt.Println(common.Version())
		} else {
			fmt.Printf("version: %s\n", common.Version())
			fmt.Printf("commit: %s\n", common.GitCommit())
			fmt.Printf("build time: %s\n", common.BuildTime())
			fmt.Printf("build user: %s\n", common.BuildUser())
		}
	},
}

// TODO: just here for testing out the log command, though it might possible to make it like db command to be part
// of go.ice's built in command for managing common config
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "test log config",
	Long:  "Test log tree printer etc.",
	Run: func(cmd *cobra.Command, args []string) {
		log.PrintTree()
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start IceHub daemon",
	Long:  "Start IceHub daemon with HTTP and gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		mustLoadConfig()
		log.Info("TODO: I need to start it ....")
		// TODO: p3 check if there is already icehubd running, by port, process name etc.
		// TODO: p1 config tracer
		if err := configTracer(); err != nil {
			log.Fatal(err)
		}
		log.Info("tracer created")
		// TODO: p2 config database
		// TODO: p1 initial services (components?)
		// TODO: p1 user service, cache service etc.
		// TODO：p1 listen on port
	},
}

// TODO: check config file using gommon config
func loadConfig() error {
	if !cfgLoaded {
		// TODO: have a config reader struct instead of using static package level method
		// TODO: config file also specify logging (which package to log etc.)
		if err := config.LoadYAMLAsStruct(cfgFile, &cfg); err != nil {
			return errors.WithMessage(err, "can't load config file")
		}
		cfgLoaded = true
	}
	return nil
}

func mustLoadConfig() {
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}
}

// FIXME: hacky function to play with tracing libraries
// https://github.com/jaegertracing/jaeger/blob/master/examples/hotrod/pkg/tracing/init.go#L31
func configTracer() error {
	tcfg := jgconfig.Configuration{
		Sampler: &jgconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jgconfig.ReporterConfig{
			LogSpans: false, // TODO: when true, enables LoggingReporter that runs in parallel with the main reporter
			// and logs all submitted spans. Main Configuration.Logger must be initialized in the code
			// for this option to have any effect.
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "localhost:6831",
		},
	}
	// TODO: a better way to use gommon/log, current tree level hierarchy may not be enough ...
	// TODO: the jaeger.Logger interface is so strange, Error(string) instead of Error(string, args ...interface{})
	// jgconfig.Logger(log)
	// TODO: Observer can be registered with the Tracer to receive notifications about new Spans.
	var err error
	tracer, closer, err = tcfg.New("service-a")
	if err != nil {
		return errors.Wrap(err, "can't create jaeger tracer")
	}
	return nil
}

func main() {
	// TODO: common root command should be put into a struct, but need another struct to store the flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "icehub.yml", "config file location")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output and set log level to debug")
	rootCmd.AddCommand(versionCmd)
	dbc := idbcmd.NewCommand(func() (icfg.DatabaseManagerConfig, error) {
		if err := loadConfig(); err != nil {
			return cfg.DatabaseManager, err
		}
		return cfg.DatabaseManager, nil
	})
	rootCmd.AddCommand(dbc.Root())
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(startCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
