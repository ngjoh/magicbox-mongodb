/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/koksmat-com/koksmat/db"
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
		log.Println("Importing")
		data := io.Readfile[model.RecipientsType](inputFile)
		db.Save[model.RecipientsType](domain, subject, data)
		log.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVarP(&inputFile, "inputFile", "i", "", "Input file (required)")
	importCmd.MarkFlagRequired("inputFile")
	importCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain (required)")
	importCmd.MarkFlagRequired("domain")
	importCmd.Flags().StringVarP(&subject, "subject", "s", "", "Subject (required)")
	importCmd.MarkFlagRequired("subject")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
