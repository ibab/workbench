package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List all available images of the current user",
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen).SprintFunc()

		images := GetImages()
		fmt.Printf("## %s: %d\n", green("images"), len(images))
		for i, image := range images {
			fmt.Printf("[%d] %s\n", i+1, *image.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(imagesCmd)
}
