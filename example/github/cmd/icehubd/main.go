package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	jgconfig "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"

	icli "github.com/at15/go.ice/ice/cli"
	icfg "github.com/at15/go.ice/ice/config"
	idbcmd "github.com/at15/go.ice/ice/db/cmd"
	igrpc "github.com/at15/go.ice/ice/transport/grpc"

	"github.com/at15/go.ice/example/github/pkg/common"
	mygrpc "github.com/at15/go.ice/example/github/pkg/transport/grpc"
	"github.com/at15/go.ice/example/github/pkg/util/logutil"

	_ "github.com/at15/go.ice/ice/db/adapters/mysql"
	_ "github.com/at15/go.ice/ice/db/adapters/postgres"
	_ "github.com/at15/go.ice/ice/db/adapters/sqlite"
)

const (
	myname = "icehubd" // 你的名字
)

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}

var cli *icli.Root
var log = logutil.Registry

// global configuration instance
var cfg common.ServerConfig

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
		useHttp := true
		useGrpc := false
		// start hg
		// start g
		// start h
		if len(args) > 0 {
			if strings.Contains(args[0], "h") {
				useHttp = true
			}
			if strings.Contains(args[0], "g") {
				useGrpc = true
			}
		}
		// TODO: need two go routine if we want to start two server
		if useHttp {
			log.Info("TODO: start http server")
		}
		if useGrpc {
			log.Info("start grpc server")
			srv, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
				mygrpc.RegisterIceHubServer(s, mygrpc.NewServer())
			})
			if err != nil {
				log.Fatalf("can't create grpc server %v", err)
			}
			if err := srv.Run(); err != nil {
				log.Fatalf("can't run grpc server %v", err)
			}
		}
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
		//registry, err := server.NewRegistry(cfg)
		//if err != nil {
		//	log.Fatalf("failed to create server registry %v", err)
		//}
		//registry.ConfigHttpHandler()
		//srv := http.NewServer(cfg.Http, registry.HTTPHandler())
		//if err := srv.Run(); err != nil {
		//	log.Fatalf("failed to start http server %v", err)
		//}
	},
}

func mustLoadConfig() {
	if err := cli.LoadConfigTo(&cfg); err != nil {
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
	cli = icli.New(
		icli.Name(myname),
		icli.Description("IceHub is an example GitHub integration service using go.ice"),
		icli.Version(buildInfo),
		icli.LogRegistry(log))
	root := cli.Command()
	dbc := idbcmd.NewCommand(func() (icfg.DatabaseManagerConfig, error) {
		if err := cli.LoadConfigTo(&cfg); err != nil {
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
