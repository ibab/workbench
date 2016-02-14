package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var terminateCmd = &cobra.Command{
	Use:   "terminate [instance]",
	Short: "Terminate an instance",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Need to specify an instance to terminate")
			os.Exit(1)
		}

		num, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		instances := GetInstances()

		if num < 1 || num > len(instances) {
			fmt.Printf("Error: No instance with index %d\n", num)
			os.Exit(1)
		}

		TerminateInstance(instances[num-1])
		fmt.Printf("Instance %s terminated.\n", *instances[num-1].PublicDnsName)

	},
}

func init() {
	RootCmd.AddCommand(terminateCmd)
}
