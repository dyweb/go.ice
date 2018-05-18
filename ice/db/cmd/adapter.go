package cmd

import (
	"fmt"

	"github.com/dyweb/go.ice/ice/db"
	"github.com/spf13/cobra"
)

var adapterCmd = &cobra.Command{
	Use:   "adapters",
	Short: "registered database adapters",
	Long:  "Show registered ice/db/adapters, not to be confused with drivers",
	Run: func(cmd *cobra.Command, args []string) {
		adapters := db.Adapters()
		if len(adapters) == 0 {
			fmt.Println("no ice/db/adapters adapter registered")
			return
		}
		fmt.Printf("%d adapters registered\n", len(adapters))
		for _, a := range adapters {
			fmt.Println(a)
		}
	},
}
