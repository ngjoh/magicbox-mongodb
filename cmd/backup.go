/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/koksmat-com/koksmat/audit"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var backupCmd = &cobra.Command{
	Use:   "backup [database]",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		switch fmt.Sprint(args[0]) {
		case "database":

			err := audit.Aggregate()
			if err != nil {
				log.Fatalln(err)
			}
		case "collection":
			log.Fatalln("NOT IMPLEMENTED")
			// err := model.SyncDomains()
			// if err != nil {
			// 	log.Fatalln(err)
			// }

		default:

			log.Fatalln("Cannot use that name", subject)
			return
		}

		log.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVarP(&hubSiteID, "hubSiteId", "", "", "Hub Site ID")

}
