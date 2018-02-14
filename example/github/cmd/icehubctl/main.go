package main

import (
	"fmt"
	"os"

	"github.com/at15/go.ice/ice"

	"github.com/at15/go.ice/example/github/pkg/common"
)

// TODO: create grpc client and try ping ....

func main() {
	app := ice.New(
		ice.Name("icehubctl"),
		ice.Description("Client of IceHub, which is an example GitHub integration service using go.ice"),
		ice.Version(common.Version()))
	root := ice.NewCmd(app)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
