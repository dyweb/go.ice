package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	icli "github.com/dyweb/go.ice/ice/cli"
	icfg "github.com/dyweb/go.ice/ice/config"
	idbcmd "github.com/dyweb/go.ice/ice/db/cmd"
	itrace "github.com/dyweb/go.ice/ice/tracing"
	igrpc "github.com/dyweb/go.ice/ice/transport/grpc"
	ihttp "github.com/dyweb/go.ice/ice/transport/http"

	"github.com/at15/go.ice/example/github/pkg/common"
	"github.com/at15/go.ice/example/github/pkg/server"
	mygrpc "github.com/at15/go.ice/example/github/pkg/transport/grpc"
	"github.com/at15/go.ice/example/github/pkg/util/logutil"

	_ "github.com/dyweb/go.ice/ice/db/adapters/mysql"
	_ "github.com/dyweb/go.ice/ice/db/adapters/postgres"
	_ "github.com/dyweb/go.ice/ice/db/adapters/sqlite"
	_ "github.com/dyweb/go.ice/ice/tracing/jaeger"
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
	Example: `
Start both grpc and http server
	icehubd start gh
`,
	Run: func(cmd *cobra.Command, args []string) {
		mustLoadConfig()
		tmgr, err := itrace.NewManager(cfg.Tracing)
		if err != nil {
			// FIXME: https://github.com/dyweb/gommon/issues/48 Fatalf source is incorrect
			log.Fatalf("can't create trace manager %v", err)
			return
		}
		tracer, err := tmgr.Tracer("icehub")
		if err != nil {
			log.Fatalf("can't create tracer %v", err)
			return
		}
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
		var wg sync.WaitGroup
		wg.Add(2)
		if useHttp {
			log.Info("start http server")
			httpSrv, err := server.NewHttpServer()
			if err != nil {
				log.Fatalf("can't create http server")
			}
			httpTrans, err := ihttp.NewServer(cfg.Http, httpSrv.Handler(), tracer)
			if err != nil {
				log.Fatalf("can't create http transport %v", err)
			}
			go func() {
				if err := httpTrans.Run(); err != nil {
					wg.Done()
					log.Fatalf("can't run http server %v", err)
				}
			}()
		}
		if useGrpc {
			log.Info("start grpc server")
			grpcSrv, err := server.NewGrpcServer()
			if err != nil {
				log.Fatalf("can't create grpc server")
			}
			grpcTrans, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
				mygrpc.RegisterIceHubServer(s, grpcSrv)
			})
			if err != nil {
				log.Fatalf("can't create grpc server %v", err)
			}
			go func() {
				if err := grpcTrans.Run(); err != nil {
					wg.Done()
					log.Fatalf("can't run grpc server %v", err)
				}
			}()
		}
		wg.Wait()
		// TODO: p3 check if there is already icehubd running, by port, process name etc.
		// TODO: p2 config database
		// TODO: p1 initial services (components?)
		// TODO: p1 user service, cache service etc.
	},
}

func mustLoadConfig() {
	if err := cli.LoadConfigTo(&cfg); err != nil {
		log.Fatal(err)
	}
}

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("IceHub is an example GitHub integration service using go.ice"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
		icli.IsServer())
	root := cli.Command()
	dbc := idbcmd.New(func() (icfg.DatabaseManagerConfig, error) {
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
