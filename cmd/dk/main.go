// dk is a docker cli using go.ice's cli and dockerclient package
package main

import (
	"context"
	"os"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
	dlog "github.com/dyweb/gommon/log"
	"github.com/olekukonko/tablewriter"

	"github.com/dyweb/go.ice/lib/dockerclient"
)

// TODO: start using cobra for handling sub commands
// TODO: support output like kubectl
var log, logReg = dlog.NewApplicationLoggerAndRegistry("dk")

func main() {
	//version()
	//images()
	//containers()
	containerLog(os.Args[1])
}

func ping() {
	c := cli()
	p, err := c.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", p)
}

func version() {
	c := cli()
	p, err := c.Version()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", p)
}

func images() {
	c := cli()
	images, err := c.ImageList(context.Background(), types.ImageListOptions{
		All: false,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%d images", len(images))
}

func containers() {
	c := cli()
	containers, err := c.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infof("total %d containers", len(containers))
	// TODO: print as table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "id", "status", "image"})
	for i, c := range containers {
		table.Append([]string{
			strconv.Itoa(i),
			c.ID,
			c.Status,
			c.Image,
		})
	}
	table.Render()
}

// https://docs.docker.com/engine/reference/commandline/logs/#options
func containerLog(container string) {
	c := cli()
	res, err := c.ContainerLog(context.Background(), container, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Follow:     true,
		Tail:       "100",
		Details:    true,
	})
	if err != nil {
		log.Fatal(err)
	}
	// https://github.com/docker/cli/blob/master/cli/command/container/logs.go
	// NOTE: using StdCopy will remove the 8 bytes multiplex header ...
	// TODO: the time in container is in different timezone by default, also docker timestamp seems to be using UTC?
	// actually it's not supported https://github.com/moby/moby/issues/33778
	stdcopy.StdCopy(os.Stdout, os.Stderr, res)
	//io.Copy(os.Stdout, res)
}

func cli() *dockerclient.Client {
	c, err := dockerclient.New("/var/run/docker.sock")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}
