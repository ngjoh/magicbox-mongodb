/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/koksmat-com/koksmat/worker"
	"github.com/spf13/cobra"
)

// workCmd represents the work command
var workCmd = &cobra.Command{
	Use:   "work",
	Short: "Starts the worker process",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		worker.Work()
	},
}

func init() {
	rootCmd.AddCommand(workCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
