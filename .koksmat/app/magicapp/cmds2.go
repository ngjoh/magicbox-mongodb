package magicapp

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/365admin/kubernetes-management/cmds"
)

func RegisterCmds() {
	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "40 backup",
		Long:  `Describe the main purpose of this kitchen`,
	}
	BackupDiscoverPostCmd := &cobra.Command{
		Use:   "discover",
		Short: "Database Discovery",
		Long:  `Discover databases in the cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.BackupDiscoverPost(ctx, args)
		},
	}
	backupCmd.AddCommand(BackupDiscoverPostCmd)
	BackupAllPostCmd := &cobra.Command{
		Use:   "all",
		Short: "Backup all databases",
		Long:  `Backup all databases in the cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.BackupAllPost(ctx, body, args)
		},
	}
	backupCmd.AddCommand(BackupAllPostCmd)

	RootCmd.AddCommand(backupCmd)
	setupCmd := &cobra.Command{
		Use:   "setup",
		Short: "90 deployment",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(setupCmd)
}
