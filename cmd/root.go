package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// cmd, 默认启动GUI
var rootCmd = &cobra.Command{
	Use:   "scoopplus",
	Short: "A native scoop soft pakage manager tools, a better choise to replace scoop.",
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
	newExeDir := path.Join(cwd, "root", "shims")
	CopyFile(exePath, path.Join(newExeDir, "scoopplus.exe"))
	// Config

	// add PATH
	envPath := os.Getenv("Path")
	fmt.Println(envPath)
	if !strings.Contains(envPath, newExeDir) {
		os.Setenv("Path", envPath+";"+newExeDir)
		// setx in user env
		os.StartProcess("setx", []string{"Path", "%Path%;" + newExeDir}, nil)
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
		fulldir := path.Join(dir, folder)
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
