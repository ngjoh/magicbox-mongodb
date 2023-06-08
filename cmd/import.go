/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/io"
	"github.com/koksmat-com/koksmat/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var inputFile string
var domain string
var subject string

func readAndSave[K any]() {
	data := io.Readfile[K](inputFile)
	db.Save[K](domain, subject, data)
}

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Add a JSON file to the import queue",
	Long:  `Add a JSON file to the import queue for further processing in MongoDB`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Importing")

		switch subject {
		case "recipients":
			readAndSave[model.RecipientType]()
			model.SyncRecipients(viper.GetString("DATABASE"))
		case "rooms":
			readAndSave[model.RoomType]()
		case "sharedmailboxes":
			model.ReadSharedMailboxes(inputFile)

		default:

			log.Fatalln("Unknown subject", subject)
			return
		}

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
}
