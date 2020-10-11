package cmd

import (
	"github.com/spf13/cobra"
)

var (
	weatherCmd = &cobra.Command{
		Use: "weather",
	}
)

func Execute() error {
	return weatherCmd.Execute()
}

