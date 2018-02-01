package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/dyweb/gommon/config"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	jgconfig "github.com/uber/jaeger-client-go/config"

	"github.com/at15/go.ice/_example/github/pkg/common"
	"github.com/at15/go.ice/_example/github/pkg/server"
	"github.com/at15/go.ice/_example/github/pkg/util/logutil"
	icfg "github.com/at15/go.ice/ice/config"
	idbcmd "github.com/at15/go.ice/ice/db/cmd"

	"github.com/at15/go.ice/ice"
	_ "github.com/at15/go.ice/ice/db/adapters/mysql"
	_ "github.com/at15/go.ice/ice/db/adapters/postgres"
	_ "github.com/at15/go.ice/ice/db/adapters/sqlite"
	"github.com/at15/go.ice/ice/transport/http"
)

const (
	myname = "icehubd" // 你的名字
)

var app *ice.App
var log = logutil.Registry

// global configuration instance
var cfg server.Config

// TODO: might need a registry of application instead of scatter variables around in main
var tracer opentracing.Tracer
var closer io.Closer

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
		// TODO: postpone tracing until we have server and client ready ...
		//if err := configTracer(); err != nil {
		//	log.Fatal(err)
		//}
		//log.Info("tracer created")
		// TODO: p2 config database
		// TODO: p1 initial services (components?)
		// TODO: p1 user service, cache service etc.
		// TODO：p1 listen on port
		registry, err := server.NewRegistry(cfg)
		if err != nil {
			log.Fatalf("failed to create server registry %v", err)
		}
		registry.ConfigHttpHandler()
		srv := http.NewServer(cfg.Http, registry.HTTPHandler())
		if err := srv.Run(); err != nil {
			log.Fatalf("failed to start http server %v", err)
		}
	},
}

// TODO: check config file using gommon config
func loadConfig() error {
	if !app.IsConfigLoaded() {
		// TODO: have a config reader struct instead of using static package level method
		// TODO: config file also specify logging (which package to log etc.)
		if err := config.LoadYAMLAsStruct(app.ConfigFile(), &cfg); err != nil {
			return errors.WithMessage(err, "can't load config file")
		}
		app.SetConfigLoaded()
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
	app = ice.New(
		ice.Name(myname),
		ice.Description("IceHub is an example GitHub integration service using go.ice"),
		ice.Version(common.Version()),
		ice.LogRegistry(log))
	root := ice.NewCmd(app)
	dbc := idbcmd.NewCommand(func() (icfg.DatabaseManagerConfig, error) {
		if err := loadConfig(); err != nil {
			return cfg.DatabaseManager, err
		}
		return cfg.DatabaseManager, nil
	})
	root.AddCommand(dbc.Root())
	root.AddCommand(logCmd)
	root.AddCommand(startCmd)
	// TODO: handle signal (ctrl+c etc.)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
