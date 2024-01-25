/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/koksmat-com/koksmat/connectors"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var sailCmd = &cobra.Command{
	Use:   "sail ",
	Short: "Auto pilot mode",
	Args:  cobra.MinimumNArgs(0),
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Sailing")
		connectors.Execute("open", connectors.Options{Dir: "/Users/nielsgregersjohansen/servers/mate/apps/www"}, "http://localhost:3010")

		connectors.Execute("npm", connectors.Options{Dir: "/Users/nielsgregersjohansen/servers/mate/apps/www"}, "run", "start")

		// restapi.Sail()
		//webserver.Run()
	},
}

func init() {
	rootCmd.AddCommand(sailCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
