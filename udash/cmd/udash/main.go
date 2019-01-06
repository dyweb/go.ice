package main

import (
	"github.com/dyweb/go.ice/udash/pkg"
	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
)

var log, logReg = dlog.NewApplicationLoggerAndRegistry("udash")

func main() {
	root := cobra.Command{
		Use:          "udash",
		Short:        "Univerisal dashboard, a demo for go.ice",
		SilenceUsage: true,
	}
	root.AddCommand(srvCmd())
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}

func srvCmd() *cobra.Command {
	addr := "localhost:9331"
	cmd := cobra.Command{
		Use:     "srv",
		Aliases: []string{"serve", "start", "servers"},
		Short:   "start server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return serve(addr)
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "localhost:9331", "address:port to list on")
	return &cmd
}

func serve(addr string) error {
	srv, err := pkg.NewServer()
	if err != nil {
		return err
	}
	if err := srv.Run(addr); err != nil {
		return err
	}
	return nil
}
