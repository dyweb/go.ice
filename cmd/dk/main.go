// dk is a docker cli using go.ice's cli and dockerclient package
package main

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"

	"github.com/dyweb/go.ice/lib/dockerclient"
)

func main() {
	//version()
	images()
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

func cli() *dockerclient.Client {
	c, err := dockerclient.New("/var/run/docker.sock")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}
