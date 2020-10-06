package cmd

import "github.com/spf13/cobra"

func init() {

}

var weather = &cobra.Command{
	Use:   "weather",
	Short: "check weather tool",
	Long:  "that is a check weather tool for command line",

}
