package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/docker/docker/api/types"
	icli "github.com/dyweb/go.ice/cli"
	"github.com/dyweb/go.ice/dockerclient"
	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
)

const (
	myname = "bh"
)

var logReg = dlog.NewRegistry()
var log = logReg.Logger()

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}

var cli *icli.Root

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("BenchHub"),
		icli.Version(buildInfo),
	)
	root := cli.Command()
	psCmd := cobra.Command{
		Use: "ps",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := mustClient()
			containers, err := c.ContainerList(context.Background(), types.ContainerListOptions{
				All: true,
			})
			if err != nil {
				return err
			}
			log.Infof("%d", len(containers))
			return nil
		},
	}
	pullCmd := cobra.Command{
		Use: "pull",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := mustClient()
			reader, err := c.ImagePull(context.Background(), "dyweb/go-dev:1.13.6", types.ImagePullOptions{})
			if err != nil {
				return err
			}
			// TODO: it's actually json stream ...
			io.Copy(os.Stdout, reader)
			reader.Close()
			return nil
		},
	}
	root.AddCommand(&psCmd)
	root.AddCommand(&pullCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mustClient() *dockerclient.Client {
	c, err := dockerclient.New("/var/run/docker.sock")
	if err != nil {
		log.Fatal(err)
	}
	return c
}
