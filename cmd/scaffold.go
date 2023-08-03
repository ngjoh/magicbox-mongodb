/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/koksmat-com/koksmat/scaffold"
	"github.com/spf13/cobra"
)

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold [inputfile]",
	Short: "Generate Type Script from PNP Template",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			scaffold.Pnp2Ts(args[0])
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)

}
