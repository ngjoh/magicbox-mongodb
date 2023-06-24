/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/koksmat-com/koksmat/excel"
	"github.com/spf13/cobra"
)

var file string

// excelCmd represents the excel command
var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "Excel stuff",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		excel.Read(file)
	},
}

func init() {
	rootCmd.AddCommand(excelCmd)
	excelCmd.Flags().StringVarP(&file, "file", "f", "", "File (required)")
	excelCmd.MarkFlagRequired("file")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// excelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// excelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
