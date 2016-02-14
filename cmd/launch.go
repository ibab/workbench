package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-input"
	"os"
	"strconv"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch an instance based on an image",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Error: Need to specify an image to launch")
			os.Exit(1)
		}

		num, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		images := GetImages()

		if num < 1 || num > len(images) {
			fmt.Printf("Error: No image with index %d\n", num)
			os.Exit(1)
		}

		ui := &input.UI{
			Writer: os.Stdout,
			Reader: os.Stdin,
		}

		query := "Which instance type should be launched?"
		instanceType, err := ui.Ask(query, &input.Options{
			Default:  "c3.4xlarge",
			Required: true,
			Loop:     true,
		})
		if err != nil {
			panic(err)
		}

		names, prices := GetSpotZonesAndPrices(instanceType)
		fmt.Println("Availability zones and prices:")
		for i, name := range names {
			price := prices[i]
			fmt.Printf("%s: %s\n", *name, *price)
		}

		zone, err := ui.Ask("Which availability zone should be chosen?", &input.Options{
			Default:  "us-east-1d",
			Required: true,
			Loop:     true,
		})
		if err != nil {
			panic(err)
		}

		price, err := ui.Ask("Which price should be chosen?", &input.Options{
			Default:  "0.0",
			Required: true,
			Loop:     true,
		})
		if err != nil {
			panic(err)
		}

		LaunchSpotInstance(
			price,
			zone,
			*images[num-1].ImageId,
			instanceType,
		)

		fmt.Printf("Spot instance of type %s requested in %s for %s.\n", instanceType, zone, price)
	},
}

func init() {
	RootCmd.AddCommand(launchCmd)
}
