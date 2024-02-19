/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"
	"path/filepath"

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
			Short: "Exports HTML from Markdown in code",
			Args:  cobra.MinimumNArgs(1),
			Long:  ``,

			Run: func(cmd *cobra.Command, args []string) {

				file, _ := url.QueryUnescape(args[0])
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

				html, _, err := kitchen.ParseMarkdown(false, filepath.Dir(file), markdown)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(html)

			},
		})

}
