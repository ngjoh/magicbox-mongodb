/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/koksmat-com/koksmat/model"
	"github.com/spf13/cobra"
)

// webhookCmd represents the webhook command
var webhookCmd = &cobra.Command{
	Use:   "webhook  [argument] ",
	Short: "Work with webhooks",
	Args:  cobra.MinimumNArgs(1),
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		service := args[0]
		switch service {
		case "parse":
			//
			model.RunWebhookEventParser()

		default:

			log.Fatalln("Unknown service", service)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webhookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webhookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
