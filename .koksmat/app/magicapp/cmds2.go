package magicapp

import "github.com/spf13/cobra"

func RegisterCmds() {
	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "40 backup",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(backupCmd)
	setupCmd := &cobra.Command{
		Use:   "setup",
		Short: "90 timerjobs",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(setupCmd)
}
