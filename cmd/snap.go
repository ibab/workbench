package cmd

import (
	"github.com/spf13/cobra"
)

var snapCmd = &cobra.Command{
	Use:   "snap",
	Short: "Create a snapshot based on a running instance",
	Run: func(cmd *cobra.Command, args []string) {
		panic("Not implemented yet")
	},
}

func init() {
	RootCmd.AddCommand(snapCmd)
}
