/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/koksmat-com/koksmat/model"

	"github.com/spf13/cobra"
)

// accessCmd represents the access command
var findCmd = &cobra.Command{
	Use:   "find [SearchString]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Lookup a recipient by address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		recipient, err := model.FindRecipientByAddress(args[0])
		if err != nil {
			fmt.Println(fmt.Sprint(err))
			return
		}

		fmt.Println(fmt.Sprintf("Found: %s", recipient.DisplayName))
		fmt.Println(fmt.Sprintf("Guid: %s", recipient.Guid))
		for i, v := range recipient.EmailAddresses {
			fmt.Println(fmt.Sprintf("EmailAddresses[%d]: %s", i, v))
		}

	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
