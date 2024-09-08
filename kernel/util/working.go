package util

import (
	"flag"
	"github.com/gofrs/flock"
	"kernel/common"
	"kernel/common/file"
	"kernel/common/log"
	"kernel/conf"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/common-nighthawk/go-figure"
)

var Mode = "dev"

//var Mode = "prod"

var (
	ServerURL       *url.URL     // 内核服务 URL
	ServerPort      = "0"        // HTTP/WebSocket 端口，0 为使用随机端口
	ServerIsRunning = false      // 服务是否正在运行
	Server          *http.Server // HTTP 服务器

	ReadOnly       bool
	AccessAuthCode string
	Lang           = ""

	Container string // docker, android, ios, std
)

const (
	ContainerStd     = "std"     // 桌面端
	ContainerDocker  = "docker"  // Docker 容器端
	ContainerAndroid = "android" // Android 端
	ContainerIOS     = "ios"     // iOS 端

	LocalHost = "127.0.0.1" // 伺服地址
	FixedPort = "25566"     // 固定端口
)

var (
	WorkingDir, _ = os.Getwd()

	WorkspaceDir   = filepath.Join(WorkingDir, conf.WORKSPACE_PATH) // 工作空间目录路径
	WorkspaceLock  *flock.Flock                                     // 工作空间锁
	DataDir        string                                           // 数据目录路径
	TempDir        string                                           // 临时目录路径
	LogPath        string                                           // 配置目录下的日志文件 ReflexKey.log 路径
	DBName         = "ReflexKey.db"                                 // SQLite 数据库文件名
	DBPath         string                                           // SQLite 数据库文件路径
	AppearancePath string                                           // 配置目录下的外观目录 appearance/ 路径
	ThemesPath     string                                           // 配置目录下的外观目录下的 themes/ 路径
	IconsPath      string                                           // 配置目录下的外观目录下的 icons/ 路径

	UIProcessIDs = sync.Map{} // UI 进程 ID
)

func Boot() {

	workspacePath := flag.String("workspace", filepath.Join(WorkingDir, conf.WORKSPACE_PATH), "dir path of the workspace, default to ./workspace/")
	port := flag.String("port", FixedPort, "port of the HTTP server")
	readOnly := flag.String("readonly", "false", "read-only mode")
	accessAuthCode := flag.String("accessAuthCode", "", "access auth code")
	ssl := flag.Bool("ssl", false, "for https and wss")
	lang := flag.String("lang", "", "zh_CN/zh_CHT/en_US/fr_FR/es_ES/ja_JP")
	mode := flag.String("mode", "prod", "dev/prod")
	flag.Parse()

	if "" != *lang {
		Lang = *lang
	}
	Mode = *mode
	ServerPort = *port
	ReadOnly, _ = strconv.ParseBool(*readOnly)
	AccessAuthCode = *accessAuthCode

	initWorkspaceDir(*workspacePath)

	SSL = *ssl

	// 工作空间仅允许被一个内核进程伺服
	tryLockWorkspace()

	bootBanner := figure.NewColorFigure(conf.NAME, "isometric3", "green", true)
	log.Info("\n" + bootBanner.String())
	logBootInfo()
}

func initWorkspaceDir(workspaceArg string) {
	WorkspaceDir = workspaceArg
	DataDir = filepath.Join(WorkspaceDir, "data")
	TempDir = filepath.Join(WorkspaceDir, "temp")
	LogPath = filepath.Join(TempDir, "log", "ReflexKey.log")
	DBPath = filepath.Join(DataDir, DBName)
	log.SetLogPath(LogPath)

	if !file.IsExist(WorkspaceDir) {
		if err := os.MkdirAll(WorkspaceDir, 0755); nil != err && !os.IsExist(err) {
			log.Error("create default workspace folder [%s] failed: %s", WorkspaceDir, err)
			os.Exit(log.ExitCodeInitWorkspaceErr)
		}
	}
	log.Info("use the workspace [%s]", WorkspaceDir)

	if err := os.MkdirAll(DataDir, 0755); nil != err && !os.IsExist(err) {
		log.Fatal(log.ExitCodeInitWorkspaceErr, "create data folder [%s] failed: %s", DataDir, err)
	}
	if err := os.MkdirAll(TempDir, 0755); nil != err && !os.IsExist(err) {
		log.Fatal(log.ExitCodeInitWorkspaceErr, "create temp folder [%s] failed: %s", TempDir, err)
	}

	assets := filepath.Join(DataDir, "assets")
	if err := os.MkdirAll(assets, 0755); nil != err && !os.IsExist(err) {
		log.Fatal(log.ExitCodeInitWorkspaceErr, "create data assets folder [%s] failed: %s", assets, err)
	}

	templates := filepath.Join(DataDir, "templates")
	if err := os.MkdirAll(templates, 0755); nil != err && !os.IsExist(err) {
		log.Fatal(log.ExitCodeInitWorkspaceErr, "create data templates folder [%s] failed: %s", templates, err)
	}

	widgets := filepath.Join(DataDir, "widgets")
	if err := os.MkdirAll(widgets, 0755); nil != err && !os.IsExist(err) {
		log.Fatal(log.ExitCodeInitWorkspaceErr, "create data widgets folder [%s] failed: %s", widgets, err)
	}

	plugins := filepath.Join(DataDir, "plugins")
	if err := os.MkdirAll(plugins, 0755); nil != err && !os.IsExist(err) {
		log.Fatal(log.ExitCodeInitWorkspaceErr, "create data plugins folder [%s] failed: %s", widgets, err)
	}

	emojis := filepath.Join(DataDir, "emojis")
	if err := os.MkdirAll(emojis, 0755); nil != err && !os.IsExist(err) {
		log.Fatal(log.ExitCodeInitWorkspaceErr, "create data emojis folder [%s] failed: %s", widgets, err)
	}

	public := filepath.Join(DataDir, "public")
	if err := os.MkdirAll(public, 0755); nil != err && !os.IsExist(err) {
		log.Fatal(log.ExitCodeInitWorkspaceErr, "create data public folder [%s] failed: %s", widgets, err)
	}
}

func tryLockWorkspace() {
	WorkspaceLock = flock.New(filepath.Join(WorkspaceDir, ".lock"))
	ok, err := WorkspaceLock.TryLock()
	if ok {
		return
	}
	if nil != err {
		log.Error("lock workspace [%s] failed: %s", WorkspaceDir, err)
	} else {
		log.Error("lock workspace [%s] failed", WorkspaceDir)
	}
	os.Exit(log.ExitCodeWorkspaceLocked)
}

func IsWorkspaceLocked(workspacePath string) bool {
	if !file.IsDir(workspacePath) {
		return false
	}

	lockFilePath := filepath.Join(workspacePath, ".lock")
	if !file.IsExist(lockFilePath) {
		return false
	}

	f := flock.New(lockFilePath)
	defer func(f *flock.Flock) {
		_ = f.Unlock()
	}(f)
	ok, _ := f.TryLock()
	if ok {
		return false
	}
	return true
}

func UnlockWorkspace() {
	if nil == WorkspaceLock {
		return
	}

	if err := WorkspaceLock.Unlock(); nil != err {
		log.Error("unlock workspace [%s] failed: %s", WorkspaceDir, err)
		return
	}

	if err := os.Remove(filepath.Join(WorkspaceDir, ".lock")); nil != err {
		log.Error("remove workspace lock failed: %s", err)
		return
	}
}

func logBootInfo() {
	plat := common.GetOSPlatform()
	log.Info("kernel is booting:\n"+
		"    * ver [%s]\n"+
		"    * arch [%s]\n"+
		"    * os [%s]\n"+
		"    * pid [%d]\n"+
		"    * runtime mode [%s]\n"+
		"    * working directory [%s]\n"+
		"    * read only [%v]\n"+
		"    * container [%s]\n"+
		"    * database [ver=%s]\n"+
		"    * workspace directory [%s]",
		conf.VERSION, runtime.GOARCH, plat, os.Getpid(), Mode, WorkingDir, ReadOnly, Container, conf.DatabaseVer, WorkspaceDir)
}
