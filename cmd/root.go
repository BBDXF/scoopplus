package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

/// cmds:

/// init              // 初始化scoop目录和配置文件
/// docktor           // 检查缺少的依赖，并安装

/// list
/// install 7zip      // 支持 online
/// uninstall 7zip
/// update 7zip       // 支持 online
/// search 7zip       // 支持 online
/// info 7zip         // 支持 online
/// status            // 支持 online

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

// cmd, 默认启动GUI
var rootCmd = &cobra.Command{
	Use:     "scoopplus",
	Version: G_VERSION,
	Short:   "A native scoop soft pakage manager tools, a better choise to replace scoop.",
}

// cmd, 初始化scoop目录和配置文件
var rootCmd_init = &cobra.Command{
	Use:   "init",
	Short: "Initialize scoop dirs and config files.",
	Run: func(cmd *cobra.Command, args []string) {
		ScoopPlusInit()
	},
}

func init() {
	rootCmd.AddCommand(rootCmd_init)
	cwd, _ := os.Getwd()
	LoadConfig(cwd)
}

func Execute() error {
	return rootCmd.Execute()
}

func ScoopPlusInit() {
	cwd, _ := os.Getwd()
	var yn string
	fmt.Println("Current working directory: ", cwd)
	fmt.Println("Are you sure to initialize here? [y/n]")
	fmt.Scanln(&yn)
	if yn == "y" || yn == "Y" || yn == "yes" || yn == "Yes" {
		fmt.Println("Initializing...")
		ScoopPlusInstall(cwd)
		fmt.Println("Initialized. Please restart your terminal.")
	} else {
		fmt.Println("Canceled.")
	}
}

func ScoopPlusInstall(cwd string) {
	// rebuild folders
	ScoopFoldersBuild(cwd)
	// Copy exe
	exePath, _ := os.Executable()
	newExeDir := filepath.Join(cwd, "root", "shims")
	CopyFile(exePath, filepath.Join(newExeDir, "scoopplus.exe"))
	// Config
	SaveConfig(cwd)
	// add PATH
	envPath := os.Getenv("Path")
	fmt.Println(envPath)
	if !strings.Contains(envPath, newExeDir) {
		os.Setenv("Path", envPath+";"+newExeDir)
		// setx in user env
		exec.Command("cmd", "/C", "setx", "Path", "%Path%;"+newExeDir).Run()
	}
}

func ScoopFoldersBuild(dir string) {
	var folders = []string{
		`root\apps\scoop\current\`,
		`root\buckets\`,
		`root\cache\`,
		`root\persist\`,
		`root\shims\`,
		`root\workspace`,
		`global\`,
	}
	for _, folder := range folders {
		fulldir := filepath.Join(dir, folder)
		if _, err := os.Stat(fulldir); os.IsNotExist(err) {
			os.MkdirAll(fulldir, 0755)
			fmt.Println("Create folder: ", fulldir)
		}
	}
}

func CopyFile(src, dst string) (err error) {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()
	dstF, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstF.Close()
	_, err = io.Copy(dstF, srcF)
	return err
}
