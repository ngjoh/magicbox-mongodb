/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var sailCmd = &cobra.Command{
	Use:   "sail [journey] [id]",
	Short: "Travel management",
	Args:  cobra.MinimumNArgs(2),
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		journey := args[0]
		id := args[1]
		log.Println("Sailing", journey, id)
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
