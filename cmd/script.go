/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"

	"github.com/koksmat-com/koksmat/kitchen"
	"github.com/spf13/cobra"
)

// webhookCmd represents the webhook command
var scriptCmd = &cobra.Command{
	Use:   "script [argument] ",
	Short: "Work with scripts",
	Args:  cobra.MinimumNArgs(1),
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(scriptCmd)
	scriptCmd.AddCommand(
		&cobra.Command{
			Use:   "html [file]",
			Short: "Exports HTML from Markdown in script",
			Args:  cobra.MinimumNArgs(1),
			Long:  ``,

			Run: func(cmd *cobra.Command, args []string) {

				file, _ := url.QueryUnescape(args[0])

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
		})

}
