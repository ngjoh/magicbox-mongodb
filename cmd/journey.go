/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/koksmat-com/koksmat/journeys"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var journeyCmd = &cobra.Command{
	Use:   "journey [operation] [path]",
	Short: "Journey management",
	Args:  cobra.MinimumNArgs(2),
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		operation := args[0]
		switch operation {
		case "unpack":
			journeys.Unpack(args[1])
		case "pack":
			journeys.Pack(args[1])

		default:

			log.Fatalln("Unknown service", operation)
			return
		}
		//webserver.Run()
	},
}

func init() {
	rootCmd.AddCommand(journeyCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
