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
	Path7z    string            `json:"path_7z"`    // 7z.exe path.
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
	Url        string `json:"url"`         // 下载url.
	Hash       string `json:"hash"`        // 下载hash.
	ExtractDir string `json:"extract_dir"` // 特殊使用.
}
type JsonBucketApp struct {
	Name   string `json:"name"`   // app name. json 文件名.
	Bucket string `json:"bucket"` // bucket json path

	Version     string `json:"version"`
	Description string `json:"description"`
	License     string `json:"license"` // 可能是object, string
	Homepage    string `json:"homepage"`
	Notes       string `json:"notes"` // 可能是arrary, string

	Url        []string          `json:"url"`          // 不区分架构的下载url
	Hash       string            `json:"hash"`         // 和url对应
	ExtractDir string            `json:"extract_dir"`  // 和url对应. 支持 .zip、.7z、.tar、.gz、.lzma 和 .lzh
	ExtractTo  []string          `json:"extract_to"`   // 和url对应. 支持 .zip、.7z、.tar、.gz、.lzma 和 .lzh
	Bin        map[string]string `json:"bin"`          // 特殊使用. 不区分架构的二进制文件. 最终要使用引导程序，转换为shim
	Shortcuts  map[string]string `json:"shortcuts"`    // 特殊使用. 不区分架构的快捷方式. 最终添加到系统快捷方式中
	Persist    map[string]string `json:"persist"`      // 特殊使用. 不区分架构的持久化文件夹. 最终转换为link
	EnvSet     map[string]string `json:"env_set"`      // 特殊使用. 不区分架构的环境变量. 新建的环境变量
	EnvAddPath []string          `json:"env_add_path"` // 特殊使用. 不区分架构的环境变量. 添加到PATH中的路径
	Innosetup  bool              `json:"innosetup"`    // 是否使用innosetup.
	Installer  string            `json:"installer"`    // 特殊使用. 不区分架构的安装器. 比如: "extras/vcredist2022/vcredist2022.exe"。 object, string
	Depends    []string          `json:"depends"`      // 特殊使用. 不区分架构的依赖. 比如: ["vcredist"]。 也有可能是字符串
	Suggest    []string          `json:"suggest"`      // 特殊使用. 依赖的app提示. 比如: "vcredist": "extras/vcredist2022"。 object, string

	// PreInstall  string `json:"pre_install"`  // 特殊使用. 不区分架构的安装前脚本. 一般是一段powershell脚本. 也有可能是字符串，数组
	PostInstall string `json:"post_install"` // 特殊使用. 不区分架构的安装后脚本. 一般是一段powershell脚本. 也有可能是字符串，数组
	// Arch map[string]JsonBucketAppArch `json:"architecture"` // key: 64bit, 32bit, arm64. 与 Url 一般不同时存在。 只保留64bit，32的备用。
	// post_uninstall / pre_uninstall
	// installer / uninstaller                  // [不使用] 安装器信息的regex。 内容很乱，不使用。
	// CheckVer    object                       // [不使用] 检查版本更新信息的regex
	// AutoUpdate  object                       // [不使用] 自动更新时用的url。 内容很乱，不使用。
}
