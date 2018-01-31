package main

import (
	"fmt"
	"os"

	"github.com/at15/go.ice/ice/app"

	"github.com/at15/go.ice/_example/github/pkg/common"
)

func main() {
	a := app.New(
		app.Name("icehubctl"),
		app.Description("Client of IceHub, which is an example GitHub integration service using go.ice"),
		app.Version(common.Version()))
	root := app.NewCmd(a)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
