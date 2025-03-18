package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var G_VERSION = "v0.1.1"
var G_SCOOPPLUS_CONFIG_FILE = "scoopplus.json"
var G_MIRROR_API = "https://api.akams.cn/github"

var G_scoopplus_config ConfigScoopPlus

// load scoop json to G_scoopplus_config.
func LoadConfig(cwd string) {
	// 读取文件内容
	var file = filepath.Join(cwd, G_SCOOPPLUS_CONFIG_FILE)
	fileContent, err := os.ReadFile(file)
	if err == nil {
		// 解析 JSON 数据到 G_scoopplus_config
		err = json.Unmarshal(fileContent, &G_scoopplus_config)
		if err != nil {
			fmt.Printf("Failed to decode config file: %v\n", err)
			return
		}
	} else {
		fmt.Printf("Failed to read config file: %v\nUse default settings.\n", err)
		G_scoopplus_config = ConfigScoopPlus{
			Online: true,
			Clean:  true,
		}
	}

	// fix something
	if G_scoopplus_config.CleanList == nil {
		G_scoopplus_config.CleanList = []string{}
	}
	if G_scoopplus_config.ScoopConf == nil {
		G_scoopplus_config.ScoopConf = map[string]string{}
	}
	if G_scoopplus_config.Buckets == nil {
		G_scoopplus_config.Buckets = []ConfigBucket{
			{Name: "main", Url: "https://github.com/ScoopInstaller/Main"},
			{Name: "extras", Url: "https://github.com/ScoopInstaller/Extras"},
			{Name: "versions", Url: "https://github.com/ScoopInstaller/Versions"},
			{Name: "nonportable", Url: "https://github.com/ScoopInstaller/Nonportable"},
			{Name: "sysinternals", Url: "https://github.com/niheaven/scoop-sysinternals"},
			{Name: "nirsoft", Url: "https://github.com/ScoopInstaller/Nirsoft"},
		}
	}
	if G_scoopplus_config.Apps == nil {
		G_scoopplus_config.Apps = []ConfigApp{}
	}
	if G_scoopplus_config.Online && G_scoopplus_config.Mirror == "" {
		if len(G_scoopplus_config.Mirrors) > 0 {
			G_scoopplus_config.Mirror = G_scoopplus_config.Mirrors[0].Url
		} else {
			G_scoopplus_config.Mirror = "https://gh-proxy.net/"
			G_scoopplus_config.Mirrors = []ConfigMirror{}
		}
	}
	if _, ok := G_scoopplus_config.ScoopConf["scoop_repo"]; !ok {
		G_scoopplus_config.ScoopConf["scoop_repo"] = "https://github.com/ScoopInstaller/Scoop"
	}
	// if _, ok := G_scoopplus_config.ScoopConf["aria2-enabled"]; !ok {
	G_scoopplus_config.ScoopConf["aria2-enabled"] = "false" // disable it always
	// }
}

func SaveConfig(cwd string) {
	// 读取文件内容
	var file = filepath.Join(cwd, G_SCOOPPLUS_CONFIG_FILE)
	fileContent, err := json.MarshalIndent(G_scoopplus_config, "", "  ")
	if err != nil {
		fmt.Printf("Failed to encode config file: %v\n", err)
		return
	}
	// 写入文件内容
	err = os.WriteFile(file, fileContent, 0644)
	if err != nil {
		fmt.Printf("Failed to write config file: %v\n", err)
		return
	}
}
