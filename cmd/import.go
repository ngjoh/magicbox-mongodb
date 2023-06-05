/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/koksmat-com/koksmat/io"
	model "github.com/koksmat-com/koksmat/model/exchange"
	"github.com/spf13/cobra"
)

var inputFile string
var domain string
var subject string

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Add a JSON file to the import queue",
	Long:  `Add a JSON file to the import queue for further processing `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Importing")
		_ = io.Readfile[model.RecipientsType]("adsfdsf")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	rootCmd.Flags().StringVarP(&inputFile, "inputFile", "i", "", "Input file (required)")
	rootCmd.MarkFlagRequired("inputFile")
	rootCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain (required)")
	rootCmd.MarkFlagRequired("domain")
	rootCmd.Flags().StringVarP(&subject, "subject", "s", "", "Subject (required)")
	rootCmd.MarkFlagRequired("subject")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
