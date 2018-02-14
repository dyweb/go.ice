package main

import (
	"fmt"
	"os"

	"github.com/at15/go.ice/ice"
	"google.golang.org/grpc"
	dlog "github.com/dyweb/gommon/log"

	"github.com/at15/go.ice/example/github/pkg/common"
	mygrpc "github.com/at15/go.ice/example/github/pkg/transport/grpc"
	"github.com/spf13/cobra"
	"context"
	"github.com/at15/go.ice/example/github/pkg/icehubpb"
)

const (
	myname = "icehubctl"
)

var log = dlog.NewApplicationLogger()
var addr = "localhost:7081"
var conn *grpc.ClientConn
var client mygrpc.IceHubClient

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping server",
	Long:  "Ping server using gRPC",
	Run: func(cmd *cobra.Command, args []string) {
		mustCreateClient()
		if res, err := client.Ping(context.Background(), &icehubpb.Ping{Name: myname}); err != nil {
			log.Fatalf("failed to ping %v", err)
		} else {
			log.Infof("ping finished name is %s", res.Name)
		}
	},
}

func main() {
	app := ice.New(
		ice.Name(myname),
		ice.Description("Client of IceHub, which is an example GitHub integration service using go.ice"),
		ice.Version(common.Version()),
		ice.LogRegistry(log))
	root := ice.NewCmd(app)
	root.AddCommand(pingCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mustCreateClient() {
	var err error
	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't dial %v", err)
	}
	client = mygrpc.NewClient(conn)
}
