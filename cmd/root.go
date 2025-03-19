package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
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

var rootCmd_test = &cobra.Command{
	Use:   "test",
	Short: "Test command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test command.")
		var js = "buckets\\main\\bucket\\7zip.json"
		var scoop_root = os.Getenv("SCOOP")
		var f = filepath.Join(scoop_root, js)
		app := ScoopAppParse(f)
		dt, _ := json.MarshalIndent(app, "", "  ")
		fmt.Print(string(dt))
	},
}

func init() {
	rootCmd.AddCommand(rootCmd_init)
	rootCmd.AddCommand(rootCmd_test)
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
	envPath := os.Getenv("PATH")
	fmt.Println(envPath)
	if !strings.Contains(envPath, newExeDir) {
		os.Setenv("PATH", envPath+";"+newExeDir)
		// setx in user env
		exec.Command("cmd", "/C", "setx", "PATH", "%PATH%;"+newExeDir).Run()
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

func WhereExePath(name string) string {
	path, err := exec.LookPath(name)
	if err != nil {
		return ""
	} else {
		return path
	}
}

func Http_Get_Content(url string) []byte {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82")
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return data
}

func ScoopAppParse(json_path string) (app *JsonBucketApp) {
	var content []byte
	if strings.HasPrefix(json_path, "http") {
		content = Http_Get_Content(json_path)
	} else {
		content, _ = os.ReadFile(json_path)
	}
	result := gjson.Parse(string(content))
	if !result.Exists() {
		return nil
	}
	app = &JsonBucketApp{
		Name:        getFileNameWithoutExt(json_path),
		Bucket:      json_path,
		Version:     result.Get("version").String(),
		Description: result.Get("description").String(),
		Homepage:    result.Get("homepage").String(),
	}
	// license
	var license = result.Get("license")
	if license.IsObject() {
		app.License = license.Get("identifier").String() // 忽略 url
	} else {
		app.License = license.String()
	}

	// note
	var note = result.Get("notes")
	if note.IsArray() {
		for _, note := range note.Array() {
			app.Notes += "\n" + note.String()
		}
	} else {
		app.Notes = note.String()
	}

	// url
	var url = result.Get("url")
	var arch = result.Get("architecture")

	app.Innosetup = result.Get("innosetup").Bool()
	app.ExtractDir = result.Get("extract_dir").String()
	app.ExtractTo = _scoop_app_str_or_array(result.Get("extract_to"))
	app.Bin = _scoop_app_json_bin(result.Get("bin"))
	app.Shortcuts = _scoop_app_json_bin(result.Get("shortcuts"))
	app.Persist = _scoop_app_json_bin(result.Get("persist"))

	if url.Exists() {
		app.Url = _scoop_app_str_or_array(url)
		app.Hash = result.Get("hash").String()
	} else if arch.Exists() {
		app.Url = _scoop_app_str_or_array(arch.Get("64bit.url"))
		app.Hash = arch.Get("64bit.hash").String()

		var extract_dir = arch.Get("64bit.extract_dir")
		if extract_dir.Exists() {
			app.ExtractDir = extract_dir.String()
		}
		var extract_to = arch.Get("64bit.extract_to")
		if extract_to.Exists() {
			app.ExtractTo = _scoop_app_str_or_array(extract_to)
		}
		var bin = arch.Get("64bit.bin")
		if bin.Exists() {
			app.Bin = _scoop_app_json_bin(bin)
		}
		var shortcuts = arch.Get("64bit.shortcuts")
		if shortcuts.Exists() {
			app.Shortcuts = _scoop_app_json_bin(shortcuts)
		}
		var persist = arch.Get("64bit.persist")
		if persist.Exists() {
			app.Persist = _scoop_app_json_bin(persist)
		}
	} else {
		return nil
	}

	// depends
	app.Depends = result.Get("depends").String()
	// suggest
	app.Suggest = result.Get("suggest").String()
	// installer
	var installer = result.Get("installer")
	if installer.IsArray() {
		installer.ForEach(func(key, value gjson.Result) bool {
			app.Installer += ", " + value.String()
			return true
		})
	} else {
		app.Installer = installer.String()
	}
	// env_set
	app.EnvSet = make(map[string]string)
	result.Get("env_set").ForEach(func(key, value gjson.Result) bool {
		app.EnvSet[key.String()] = value.String()
		return true
	})
	// env_add_path
	app.EnvAddPath = []string{}
	result.Get("env_add_path").ForEach(func(key, value gjson.Result) bool {
		app.EnvAddPath = append(app.EnvAddPath, value.String())
		return true
	})

	return
}

func _scoop_app_str_or_array(val gjson.Result) (str []string) {
	str = []string{}
	if val.IsArray() {
		for _, item := range val.Array() {
			str = append(str, item.String())
		}
	} else {
		str = append(str, val.String())
	}
	return
}

func _scoop_bin_make(ls []string) (k, v string) {
	if len(ls) == 1 {
		return getFileNameWithoutExt(ls[0]), ls[0]
	} else if len(ls) == 2 {
		return ls[1], ls[0]
	}
	// 路径，别名，参数
	return ls[1], ls[0] + " " + strings.Join(ls[2:], " ")
}
func _scoop_app_json_bin(val gjson.Result) (bins map[string]string) {
	bins = make(map[string]string)
	if val.IsArray() {
		for _, item := range val.Array() {
			if item.IsArray() {
				var args = []string{}
				for _, arg := range item.Array() {
					args = append(args, arg.String())
				}
				k, v := _scoop_bin_make(args)
				bins[k] = v
			} else {
				k, v := _scoop_bin_make([]string{val.String()})
				bins[k] = v
			}
		}
	} else {
		k, v := _scoop_bin_make([]string{val.String()})
		bins[k] = v
	}
	return
}

func getFileNameWithoutExt(json_path string) string {
	// if strings.HasPrefix(json_path, "http") {
	// 	parsedURL, err := url.Parse(json_path)
	// 	if err != nil {
	// 		return ""
	// 	}
	// 	base := filepath.Base(parsedURL.Path)
	// 	ext := path.Ext(base)
	// 	return base[:len(base)-len(ext)]
	// } else {
	base := filepath.Base(json_path)
	ext := filepath.Ext(base)
	return base[:len(base)-len(ext)]
	// }
}
