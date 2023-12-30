/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/koksmat-com/koksmat/model"
	"github.com/spf13/cobra"
)

var Host string

// pwshCmd represents the pwsh command
var pwshCmd = &cobra.Command{
	Use:   "pwsh file",
	Short: "Run PowerShell",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		data, err := os.ReadFile(args[0])
		if err != nil {
			panic(err)
		}

		host := Host
		script := string(data)

		finalScript := fmt.Sprintf(`
%s
ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM

		`, script)
		result, err := model.ExecutePowerShellScript("koksmat", host, finalScript, "")
		if err != nil {
			panic(err)
		}
		println(result)

	},
}

func init() {
	rootCmd.AddCommand(pwshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	pwshCmd.Flags().StringVarP(&Host, "host", "", "", "Host system to connect to. The default is none. Options are: exchange sharepoint powerapps")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pwshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
