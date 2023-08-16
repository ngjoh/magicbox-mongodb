/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/koksmat-com/koksmat/powershell"
	"github.com/spf13/cobra"
)

var sharepointCmd = &cobra.Command{
	Use:   "sharepoint  [argument] [inputfile]",
	Short: "Work with PNP Templates and more",
	Args:  cobra.MinimumNArgs(2),
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		service := args[0]
		switch service {
		case "export-template":
			powershell.GetSiteTemplate(args[1])

		default:

			log.Fatalln("Unknown service", service)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(sharepointCmd)

}
