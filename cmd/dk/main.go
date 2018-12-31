// dk is a docker cli using go.ice's cli and dockerclient package
package main

import (
	"context"
	"os"
	"strconv"

	"github.com/docker/docker/api/types"
	dlog "github.com/dyweb/gommon/log"
	"github.com/olekukonko/tablewriter"

	"github.com/dyweb/go.ice/lib/dockerclient"
)

// TODO: start using cobra for handling sub commands
var log, logReg = dlog.NewApplicationLoggerAndRegistry("dk")

func main() {
	//version()
	//images()
	containers()
}

func ping() {
	c, err := dockerclient.New("/var/run/docker.sock")
	if err != nil {
		log.Fatal(err)
	}
	p, err := c.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", p)
}

func version() {
	c, err := dockerclient.New("/var/run/docker.sock")
	if err != nil {
		log.Fatal(err)
	}
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

func cli() *dockerclient.Client {
	c, err := dockerclient.New("/var/run/docker.sock")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}
