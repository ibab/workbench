package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen).SprintFunc()

		images := GetImages()
		fmt.Printf("## %s: %d\n", green("images"), len(images))
		for i, image :=  range images {
			fmt.Printf("[%d] %s\n", i + 1, *image.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(imagesCmd)
}
