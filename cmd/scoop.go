package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
