package cmd

/// scoopplus config json

type ConfigApp struct {
	Name        string `json:"name"`        // app name.
	Version     string `json:"version"`     // app version.
	Arch        string `json:"arch"`        // x86_64, arm64
	Url         string `json:"url"`         // app json url.
	Bucket      string `json:"bucket"`      // bucket name. online is null.
	License     string `json:"license"`     // app license.
	Homepage    string `json:"homepage"`    // app homepage url.
	Description string `json:"description"` // app description.
	UpdateTime  string `json:"update_time"` // last update time.
}

type ConfigBucket struct {
	Name       string `json:"name"`        // bucket name.
	Url        string `json:"url"`         // bucket url.
	UpdateTime string `json:"update_time"` // last update time.
}

type ConfigMirror struct {
	Url      string  `json:"url"`      // mirror url.
	Server   string  `json:"server"`   // mirror server name.
	Ip       string  `json:"ip"`       // mirror ip.
	Location string  `json:"location"` // mirror location.
	Latency  int     `json:"latency"`  // mirror latency.
	Speed    float32 `json:"speed"`    // mirror speed.
}

type ConfigScoopPlus struct {
	Online    bool              `json:"online"`     // scoop.sh online mode.
	Clean     bool              `json:"clean"`      // clean url function.
	CleanList []string          `json:"clean_list"` // clean url prefix list.
	Mirror    string            `json:"mirror"`     // github proxy mirror url.
	Mirrors   []ConfigMirror    `json:"mirrors"`    // github proxy list api url.
	ScoopConf map[string]string `json:"scoop_conf"` // scoop configs.
	Buckets   []ConfigBucket    `json:"buckets"`    // scoop buckets.
	Apps      []ConfigApp       `json:"apps"`       // scoop apps.
}

/// json struct for scoop json file

// https://api.akams.cn/github
type JsonGithubMirrors struct {
	Code  int            `json:"code"`
	Data  []ConfigMirror `json:"data"`
	Msg   string         `json:"msg"`
	Total int            `json:"total"`
	Time  string         `json:"update_time"`
}

// scoop json 规范太乱了！！！
// 不能使用常规方式解析。需要动态check然后处理。
// 这个struct作为解析完成后的存储使用。
// https://github.com/ScoopInstaller/Main/blob/master/bucket/7zip.json
type JsonBucketAppArch struct {
	Url        string `json:"url"`                   // 下载url.
	Hash       string `json:"hash,omitempty"`        // 下载hash.
	ExtractDir string `json:"extract_dir,omitempty"` // 特殊使用.
}
type JsonBucketApp struct {
	Name        string      // app name. json 文件名.
	Version     string      `json:"version"`
	Description string      `json:"description"`
	License     interface{} `json:"license"` // 可能是object, string
	Homepage    string      `json:"homepage"`
	Notes       interface{} `json:"notes,omitempty"` // 可能是arrary, string

	Url        string            `json:"url,omitempty"`         // 不区分架构的下载url
	Hash       string            `json:"hash,omitempty"`        // 和url对应
	ExtractDir string            `json:"extract_dir,omitempty"` // 和url对应
	ExtractTo  string            `json:"extract_to,omitempty"`  // 和url对应
	Bin        []string          `json:"bin,omitempty"`         // 特殊使用. 不区分架构的二进制文件. 比如: 7zip.exe, 7z.exe, 7za.exe
	Shortcuts  [][]string        `json:"shortcuts,omitempty"`   // 特殊使用. 不区分架构的快捷方式. 比如: 7zip.lnk, 7z.lnk, 7za.lnk
	Persist    interface{}       `json:"persist,omitempty"`     // 特殊使用. 不区分架构的持久化文件夹. 比如: Formats, Languages, Themes。 也有可能是字符串
	EnvSet     map[string]string `json:"env_set,omitempty"`     // 特殊使用. 不区分架构的环境变量. 比如: "NSISDIR": "$dir" 需要处理通配符

	// 安装后，提示用户的操作。不主动执行
	Suggest     string `json:"suggest,omitempty"`      // 特殊使用. 依赖的app提示. 比如: "vcredist": "extras/vcredist2022"。 object, string
	PreInstall  string `json:"pre_install,omitempty"`  // 特殊使用. 不区分架构的安装前脚本. 一般是一段powershell脚本. 也有可能是字符串，数组
	PostInstall string `json:"post_install,omitempty"` // 特殊使用. 不区分架构的安装后脚本. 一般是一段powershell脚本. 也有可能是字符串，数组

	Arch map[string]JsonBucketAppArch `json:"architecture,omitempty"` // key: 64bit, 32bit, arm64. 与 Url 一般不同时存在。 只保留64bit，32的备用。

	// post_uninstall / pre_uninstall
	// installer / uninstaller                  // [不使用] 安装器信息的regex。 内容很乱，不使用。
	// CheckVer    object                       // [不使用] 检查版本更新信息的regex
	// AutoUpdate  object                       // [不使用] 自动更新时用的url。 内容很乱，不使用。
}
