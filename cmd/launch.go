package cmd

import (
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch an instance based on an image",
	Run: func(cmd *cobra.Command, args []string) {
		panic("Not implemented yet")
	},
}

func init() {
	RootCmd.AddCommand(launchCmd)
}
