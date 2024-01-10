/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"path/filepath"

	"github.com/koksmat-com/koksmat/connectors"
	"github.com/koksmat-com/koksmat/kitchen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var kitchenName string
var stationName string
var channelName string
var tenantName string = "365adm"

func UnEscape(s string) string {
	ss, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return ss

}

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

func cmd(use string, short string, long string, run func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run:   run,
	}
	cmd.Flags().StringVarP(&kitchenName, "kitchen", "k", "", "Kitchen (required)")
	cmd.MarkFlagRequired("kitchen")
	cmd.Flags().StringVarP(&stationName, "station", "s", "", "Station (required)")
	cmd.MarkFlagRequired("station")
	return cmd
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
			filename, _ := url.QueryUnescape(args[0])
			file := path.Join(root, kitchenName, stationName, filename)
			markdown := ""
			switch filepath.Ext(file) {
			case ".ps1":
				md, _, err := kitchen.ReadMarkdownFromPowerShell(file)
				if err != nil {
					fmt.Println(err)
				}
				markdown = md
			case ".go":
				md, err := kitchen.ReadMarkdownFromGo(file)
				if err != nil {
					fmt.Println(err)
				}
				markdown = md
			default:
				fmt.Println("Unknown file type")
				return
			}

			html, _, err := kitchen.ParseMarkdown(filepath.Dir(file), markdown)
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

			fileName, err := url.QueryUnescape(args[0])
			if err != nil {
				log.Fatalln(err)
			}
			metadata, err := kitchen.GetMetadata(kitchenName, stationName, fileName)
			printJSON(metadata)

		},
	}
	scriptcmd.AddCommand(metaCmd)
	metaCmd.Flags().StringVarP(&kitchenName, "kitchen", "k", "", "Kitchen (required)")
	metaCmd.MarkFlagRequired("kitchen")
	metaCmd.Flags().StringVarP(&stationName, "station", "s", "", "Station (required)")
	metaCmd.MarkFlagRequired("station")

	scriptcmd.AddCommand(cmd("edit [file]", "Edit script", "", func(cmd *cobra.Command, args []string) {
		root := viper.GetString("KITCHENROOT")

		filename, _ := url.QueryUnescape(args[0])

		file := path.Join(root, kitchenName, stationName, filename)
		connectors.Execute("code", *&connectors.Options{}, file)
		fmt.Println("Opened", file)

	}))

	runcmd := cmd("run [file]", "Run script", "", func(cmd *cobra.Command, args []string) {

		filename := UnEscape(args[0])
		mateContext, err := connectors.GetContext()
		if err != nil {
			log.Fatalln(err)
		}
		result, err := kitchen.Cook(mateContext.Tenant, kitchenName, stationName, filename, nil)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(result)

	})

	runcmd.Flags().StringVarP(&channelName, "channel", "c", "", "Centrifugo channel to write back on")
	scriptcmd.AddCommand(runcmd)

}
