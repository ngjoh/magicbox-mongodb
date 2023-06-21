/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/koksmat-com/koksmat/model"
	"github.com/koksmat-com/koksmat/restapi"
	"github.com/spf13/cobra"
)

var identity string

// accessCmd represents the access command
var accessCmd = &cobra.Command{
	Use:   "access [identity]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Add or update the key of a given access identity",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		colorReset := "\033[0m"
		colorYellow := "\033[33m"
		id := args[0]
		key, _, _ := model.IssueAccessKey(id)

		token, _ := restapi.IssueIdToken(id, key)

		fmt.Println(colorYellow + token + colorReset)
		clipboard.WriteAll(token)
		fmt.Println("Copied to clipboard")
	},
}

func init() {
	rootCmd.AddCommand(accessCmd)
	accessCmd.Flags().StringVarP(&identity, "identity", "i", "", "Input file (required)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
