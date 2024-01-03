/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"path"

	"github.com/koksmat-com/koksmat/kitchen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var kitchenName string
var stationName string

// serveCmd represents the serve command
var kitchenCmd = &cobra.Command{
	Use:   "kitchen [[service]]",
	Short: "kitchen  ",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No service specified")
			return
		}
		service := args[0]
		switch service {
		case "list":
			k, err := kitchen.List()
			if err != nil {
				log.Fatalln(err)
			}
			printJSON(k)
			// restapi.All()

		default:

			log.Fatalln("Unknown service", service)
			return
		}
		//webserver.Run()
	},
}

var scriptcmd = &cobra.Command{
	Use:   "script [script]",
	Short: "Working with scripts",
	Long:  ``,
}

func init() {

	rootCmd.AddCommand(kitchenCmd)
	kitchenCmd.AddCommand(

		&cobra.Command{
			Use:   "stations [kitchen]",
			Short: "List stations in kitchen",
			Args:  cobra.MinimumNArgs(1),
			Long:  ``,

			Run: func(cmd *cobra.Command, args []string) {
				name := args[0]
				stations, err := kitchen.GetStations(name)
				if err != nil {
					log.Fatalln(err)
				}
				printJSON(stations)

				// kitchen := args[0]

			},
		})

	kitchenCmd.AddCommand(

		&cobra.Command{
			Use:   "status [kitchen]",
			Short: "Get status of kitchen",
			Args:  cobra.MinimumNArgs(1),
			Long:  ``,

			Run: func(cmd *cobra.Command, args []string) {
				name := args[0]
				status, err := kitchen.GetStatus(name)
				if err != nil {
					log.Fatalln(err)
				}
				printJSON(status)

				// kitchen := args[0]

			},
		})

	kitchenCmd.AddCommand(scriptcmd)

	htmlCmd := &cobra.Command{
		Use:   "html [file]",
		Short: "Exports HTML from Markdown in script",
		Args:  cobra.MinimumNArgs(1),
		Long:  ``,

		Run: func(cmd *cobra.Command, args []string) {
			root := viper.GetString("KITCHENROOT")
			filename := args[0]
			file := path.Join(root, kitchenName, stationName, filename)

			markdown, err := kitchen.ReadMarkdownFromPowerShell(file)
			if err != nil {
				fmt.Println(err)
			}
			html, _, err := kitchen.ParseMarkdown(markdown)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(html)

		},
	}
	scriptcmd.AddCommand(htmlCmd)
	htmlCmd.Flags().StringVarP(&kitchenName, "kitchen", "k", "", "Kitchen (required)")
	htmlCmd.MarkFlagRequired("kitchen")
	htmlCmd.Flags().StringVarP(&stationName, "station", "s", "", "Station (required)")
	htmlCmd.MarkFlagRequired("station")

	metaCmd := &cobra.Command{
		Use:   "meta [file]",
		Short: "Exports Metadata from Markdown in script",
		Args:  cobra.MinimumNArgs(1),
		Long:  ``,

		Run: func(cmd *cobra.Command, args []string) {
			type Meta struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}
			root := viper.GetString("KITCHENROOT")
			filename := args[0]
			file := path.Join(root, kitchenName, stationName, filename)

			markdown, err := kitchen.ReadMarkdownFromPowerShell(file)
			if err != nil {
				fmt.Println(err)
			}
			_, meta, err := kitchen.ParseMarkdown(markdown)
			if err != nil {
				fmt.Println(err)
			}

			metadata := Meta{
				Title:       kitchen.GetMetadataProperty(meta, "title", ""),
				Description: kitchen.GetMetadataProperty(meta, "description", ""),
			}
			printJSON(metadata)

		},
	}
	scriptcmd.AddCommand(metaCmd)
	metaCmd.Flags().StringVarP(&kitchenName, "kitchen", "k", "", "Kitchen (required)")
	metaCmd.MarkFlagRequired("kitchen")
	metaCmd.Flags().StringVarP(&stationName, "station", "s", "", "Station (required)")
	metaCmd.MarkFlagRequired("station")

}
