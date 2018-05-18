package cli

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"os"
)

type BuildInfo struct {
	Version   string
	Commit    string
	BuildTime string
	BuildUser string
	GoVersion string
}

func (info *BuildInfo) PrintTo(w io.Writer) {
	fmt.Fprintf(w, "version: %s\n", info.Version)
	fmt.Fprintf(w, "commit: %s\n", info.Commit)
	fmt.Fprintf(w, "build time: %s\n", info.BuildTime)
	fmt.Fprintf(w, "build user: %s\n", info.BuildUser)
	fmt.Fprintf(w, "go version: %s\n", info.GoVersion)
}

func makeVersionCmd(root *root) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "print version",
		Long:  "Print current version " + root.Version(),
		Run: func(cmd *cobra.Command, args []string) {
			if root.verbose {
				root.buildInfo.PrintTo(os.Stdout)
			} else {
				fmt.Println(root.Version())
			}
		},
	}
}
