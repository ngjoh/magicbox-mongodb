package magicapp

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/365admin/magicbox-mongodb/cmds"
)

func RegisterCmds() {
	discoverCmd := &cobra.Command{
		Use:   "discover",
		Short: "Disover",
		Long:  `Describe the main purpose of this kitchen`,
	}
	DiscoverDiscoverPostCmd := &cobra.Command{
		Use:   "list",
		Short: "Database Discovery",
		Long:  `Discover databases in the cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			fmt.Print(cmds.DiscoverDiscoverPost(ctx, args))
		},
	}
	discoverCmd.AddCommand(DiscoverDiscoverPostCmd)

	RootCmd.AddCommand(discoverCmd)
	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "Backup",
		Long:  `Describe the main purpose of this kitchen`,
	}
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
	restoreCmd := &cobra.Command{
		Use:   "restore",
		Short: "Backup",
		Long:  `Describe the main purpose of this kitchen`,
	}
	RestoreListPostCmd := &cobra.Command{
		Use:   "list",
		Short: "List backup blobs",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.RestoreListPost(ctx, args)
		},
	}
	restoreCmd.AddCommand(RestoreListPostCmd)
	RestoreDownloadPostCmd := &cobra.Command{
		Use:   "download",
		Short: "Download all backups",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.RestoreDownloadPost(ctx, body, args)
		},
	}
	restoreCmd.AddCommand(RestoreDownloadPostCmd)
	RestoreUnarchivePostCmd := &cobra.Command{
		Use:   "unarchive",
		Short: "Database Restore",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.RestoreUnarchivePost(ctx, args)
		},
	}
	restoreCmd.AddCommand(RestoreUnarchivePostCmd)
	RestoreListtarPostCmd := &cobra.Command{
		Use:   "listtar",
		Short: "Database Restore",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.RestoreListtarPost(ctx, args)
		},
	}
	restoreCmd.AddCommand(RestoreListtarPostCmd)

	RootCmd.AddCommand(restoreCmd)
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "MongoDB",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(deployCmd)
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "MongoDB",
		Long:  `Describe the main purpose of this kitchen`,
	}

	RootCmd.AddCommand(installCmd)
}
