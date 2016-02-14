package cmd

import (
	"fmt"
	"os"
	"strconv"
	"os/exec"
	"github.com/spf13/cobra"
)

// attachCmd represents the attach command
var attachCmd = &cobra.Command{
	Use:   "attach [instance]",
	Short: "Attach to a running instance",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the instance index selected by the user
		var selected_number int
		var err error
		if len(args) > 0 {
			selected_number, err = strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			selected_number = 1
		}

		instances := GetInstances()

		if len(instances) == 0 {
			fmt.Println("Can't attach, as there are no running instances")
		} else {
			// Launch or attach tmux on the server

			if selected_number < 1 || selected_number > len(instances) {
				fmt.Printf("No such instance: %d\n", selected_number)
				os.Exit(1)
			}

			inst := instances[selected_number - 1]
			launchTmux := "'tmux has-session -t work > /dev/null 2>&1; " + 
			                  "if [ $? -eq 0 ];" + 
												"  then tmux attach-session -t work;" + 
												"  else tmux new-session -s work;" + 
												"fi'"
			cmd := exec.Command("ssh", "-t", *inst.PublicDnsName, "bash", "-c", launchTmux)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(attachCmd)
}
