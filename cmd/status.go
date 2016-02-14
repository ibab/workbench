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
		yellow := color.New(color.FgYellow).SprintFunc()

		instances := GetInstances()

		fmt.Printf("## instances: %d\n", len(instances))

		if len(instances) == 0 {
			fmt.Println("No instances are currently running")
		}

		for i, inst := range instances {
			marker := "?"
			if *inst.State.Name == "running" {
				marker = green("R")
			} else if *inst.State.Name == "pending" {
				marker = yellow("P")
			}

			fmt.Printf("[%d] %s %s\n", i+1, marker, *inst.PublicDnsName)
		}
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
