/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/koksmat-com/koksmat/powershell"
	"github.com/spf13/cobra"
)

// pwshCmd represents the pwsh command
var pwshCmd = &cobra.Command{
	Use:   "powershell",
	Short: "Run PowerShell",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		type NewSharedMailboxResult struct {
			Name               string `json:"Name"`
			DisplayName        string `json:"DisplayName"`
			Identity           string `json:"Identity"`
			PrimarySmtpAddress string `json:"PrimarySmtpAddress"`
		}
		id := fmt.Sprintf("%s", uuid.New())

		result, err := powershell.CreateSharedMailbox(id, id, id, []string{"s", "s"}, []string{"s", "s"}, []string{"s", "s"})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("DisplayName", result.DisplayName)
			fmt.Println("PrimarySmtpAddress", result.PrimarySmtpAddress)
			fmt.Println("ExchangeObjectId", result.ExchangeObjectId)
			fmt.Println("Name", result.Name)

			fmt.Println("--------------------------------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(pwshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pwshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pwshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
