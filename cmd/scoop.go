package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/// scoop cmd
/// list
/// install 7zip      // 支持 online
/// uninstall 7zip
/// update 7zip       // 支持 online

/// search 7zip       // 支持 online
/// info 7zip         // 支持 online
/// check             // 支持 online
/// cache clean

/// bucket list
/// bucket update     // update buckets and index apps
/// bucket add <bucket> <url>
/// bucket remove <bucket>

/// config <key> <value>
/// config <key>
/// config list

/// mirror <url>     // config github mirror url
/// mirror list      // list github mirror urls

/// -o --online      // use online mode, default is offline mode. 可以config online true 修改首选模式, 在线查不到才使用本地bucket.

func init() {
	scoopCmd.AddCommand(scoopCmd_install)
	rootCmd.AddCommand(scoopCmd)
}

var scoopCmd = &cobra.Command{
	Use:   "scoop",
	Short: "scoop folders re-build and install",
}

var scoopCmd_install = &cobra.Command{
	Use:   "install",
	Short: "scoop install",
	Long:  `scoop install, so you can use scoop both with this tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		scoopInstall()
	},
}

func scoopInstall() {
	fmt.Printf("TODO scoop install\r\n")
}
