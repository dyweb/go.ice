// dk is a docker cli using go.ice's cli and dockerclient package
package main

import (
	"log"

	"github.com/dyweb/go.ice/lib/dockerclient"
)

func main() {
	version()
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
