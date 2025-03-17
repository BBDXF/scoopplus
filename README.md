# ScoopPlus
ScoopPlus is a command-line tool that extends the functionality of the Scoop package manager. It provides additional commands to manage packages, including installation, removal, and updating.

# Features
- support online mode
- support mirror proxy
- support download url hook/mirror
- support GUI and CLI
- Golang implementation not powershell scripts
- Command line support compatibility with Scoop
- Disable the automatic update of buckets

```bash
# scoop cmd
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

```

# Usage
## build
```bash
# install wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# clone
git clone xxxx

# install frontend dependencies
cd frontend
npm install

# build
cd ..
wails build -windowsconsole

# development
wails dev
```

## Run GUI
使用vbs脚本隐藏cmd窗口运行
```bash
# run.vbs
set ws=WScript.CreateObject("WScript.Shell")
ws.Run "cmd /c scoopplus.exe gui",0
```

## Run CLI
```bash
scoopp.exe 
scoopp.exe -h
```


# Roadmap
## V0.1.0
- Basic framework 
- Basic GUI
- Basic CLI



# About
BBDXF
