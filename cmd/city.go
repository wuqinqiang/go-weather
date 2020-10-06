package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	weatherCmd.AddCommand(cityCmd)
}

var cityCmd = &cobra.Command{
	Use:   "city",
	Short: "city code",

}
