package model

import (
	"golang.org/x/mod/semver"
	"kernel/common"
	"kernel/conf"
	"kernel/util"
	"sync"
)

var Conf *AppConf

// AppConf 维护应用元数据，保存在 ~/.ReflexKey/conf.json。
type AppConf struct {
	LogLevel string `json:"logLevel"` // 日志级别：Off, Trace, Debug, Info, Warn, Error, Fatal
	//Appearance     *conf.Appearance `json:"appearance"`     // 外观
	//Langs          []*conf.Lang     `json:"langs"`          // 界面语言列表
	Lang string `json:"lang"` // 选择的界面语言，同 Appearance.Lang
	//FileTree       *conf.FileTree   `json:"fileTree"`       // 文档面板
	//Tag            *conf.Tag        `json:"tag"`            // 标签面板
	//Editor         *conf.Editor     `json:"editor"`         // 编辑器配置
	//Export         *conf.Export     `json:"export"`         // 导出配置
	//Graph          *conf.Graph      `json:"graph"`          // 关系图配置
	//UILayout       *conf.UILayout   `json:"uiLayout"`       // 界面布局。不要直接使用，使用 GetUILayout() 和 SetUILayout() 方法
	UserData string `json:"userData"` // 社区用户信息，对 User 加密存储
	//User           *conf.User       `json:"-"`              // 社区用户内存结构，不持久化。不要直接使用，使用 GetUser() 和 SetUser() 方法
	//Account        *conf.Account    `json:"account"`        // 帐号配置
	ReadOnly       bool         `json:"readonly"`       // 是否是以只读模式运行
	LocalIPs       []string     `json:"localIPs"`       // 本地 IP 列表
	AccessAuthCode string       `json:"accessAuthCode"` // 访问授权码
	System         *conf.System `json:"system"`         // 系统配置
	//Keymap         *conf.Keymap     `json:"keymap"`         // 快捷键配置
	//Sync           *conf.Sync       `json:"sync"`           // 同步配置
	//Search         *conf.Search     `json:"search"`         // 搜索配置
	//Flashcard      *conf.Flashcard  `json:"flashcard"`      // 闪卡配置
	//AI             *conf.AI         `json:"ai"`             // 人工智能配置
	//Bazaar         *conf.Bazaar     `json:"bazaar"`         // 集市配置
	//Stat           *conf.Stat       `json:"stat"`           // 统计
	//Api            *conf.API        `json:"api"`            // API
	//Repo           *conf.Repo       `json:"repo"`           // 数据仓库
	//Publish        *conf.Publish    `json:"publish"`        // 发布服务
	OpenHelp      bool `json:"openHelp"`      // 启动后是否需要打开用户指南
	ShowChangelog bool `json:"showChangelog"` // 是否显示版本更新日志
	CloudRegion   int  `json:"cloudRegion"`   // 云端区域，0：中国大陆，1：北美
	//Snippet        *conf.Snpt       `json:"snippet"`        // 代码片段
	DataIndexState int `json:"dataIndexState"` // 数据索引状态，0：已索引，1：未索引

	m *sync.Mutex
}

func InitConf() {

	Conf = &AppConf{LogLevel: "trace", m: &sync.Mutex{}}

	if nil == Conf.System {
		Conf.System = conf.NewSystem()
		if util.ContainerIOS != util.Container {
			Conf.OpenHelp = true
		}
	} else {
		if 0 < semver.Compare("v"+conf.VERSION, "v"+Conf.System.KernelVersion) {
			common.Info("upgraded from version [%s] to [%s]", Conf.System.KernelVersion, conf.VERSION)
			Conf.ShowChangelog = true
		} else if 0 > semver.Compare("v"+conf.VERSION, "v"+Conf.System.KernelVersion) {
			common.Info("downgraded from version [%s] to [%s]", Conf.System.KernelVersion, conf.VERSION)
		}

		Conf.System.KernelVersion = conf.VERSION
	}

}
