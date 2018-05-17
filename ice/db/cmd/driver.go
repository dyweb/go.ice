package cmd

import (
	"fmt"

	"github.com/dyweb/go.ice/ice/db"
	"github.com/spf13/cobra"
)

var driverCmd = &cobra.Command{
	Use:   "drivers",
	Short: "registered database drivers",
	Long:  "Show registered database drivers",
	Run: func(cmd *cobra.Command, args []string) {
		drivers := db.Drivers()
		if len(drivers) == 0 {
			fmt.Println("not database/sql driver registered")
			return
		}
		fmt.Printf("%d drivers registered\n", len(drivers))
		for _, d := range drivers {
			fmt.Println(d)
		}
	},
}
