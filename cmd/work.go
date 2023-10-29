/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/koksmat-com/koksmat/exchange"
	"github.com/spf13/cobra"
)

// workCmd represents the work command
var workCmd = &cobra.Command{
	Use:   "work [area name]",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting worker")

		switch fmt.Sprint(args[0]) {
		case "exchange":

			exchange.Work()

		default:

			log.Fatalln("Unknow area", subject)
			return
		}

		log.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(workCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
