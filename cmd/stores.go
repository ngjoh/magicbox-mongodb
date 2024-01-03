/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/koksmat-com/koksmat/connectors"
	"github.com/koksmat-com/koksmat/stores"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var storesCmd = &cobra.Command{
	Use:   "store [path]",
	Short: "store  ",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No service specified")
			return
		}
		path := args[0]
		switch path {
		case "context":
			printData(connectors.GetContext())
		case "mongodb/clusters":
			printData(stores.PerconaCRDS())
		// case "pod"
		// kubectl -n percona exec booking-mongos-0 -- df -h

		default:

			log.Fatalln("Unknown ", path)
			return
		}
		//webserver.Run()
	},
}

func init() {
	rootCmd.AddCommand(storesCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
