package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current status of workbench",
	Run: func(cmd *cobra.Command, args []string) {

		green := color.New(color.FgGreen).SprintFunc()

		instances := GetInstances()

		fmt.Printf("## %s: %d\n", green("running"), len(instances))

		if len(instances) == 0 {
			fmt.Println("No instances are currently running")
		}

		for i, inst := range instances {
			fmt.Printf("[%d] %s\n", i+1, *inst.PublicDnsName)
		}
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
