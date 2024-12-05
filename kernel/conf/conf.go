package conf

import (
	"bytes"
	"encoding/json"
	"kernel/common"
	"os"
	"path/filepath"
	"sync"
)

var Conf *AppConf

// AppConf 维护应用元数据，保存在 ./workspace/conf.json。
type AppConf struct {
	LogLevel string `json:"logLevel"` // 日志级别：Off, Trace, Debug, Info, Warn, Error, Fatal

	ReadOnly bool `json:"readonly"` // 是否是以只读模式运行

	AccessAuthCode string `json:"accessAuthCode"` // 访问授权码
	Authorize      bool   `json:"authorize"`      // 是否开启鉴权（开启网络伺服时默认开启）

	SSL             bool     `json:"ssl"` // 是否开启 HTTPS
	CertFileName    string   `json:"certFileName"`
	CertKeyFileName string   `json:"certKeyFileName"`
	CertDNSNames    []string `json:"certDNSNames"`

	NetworkServe bool          `json:"networkServe"` // 是否开启网络伺服
	NetworkProxy *NetworkProxy `json:"networkProxy"` // 网络代理

	m *sync.Mutex
}

func InitConf() {

	Conf = &AppConf{
		LogLevel: "trace", m: &sync.Mutex{},
	}

	confPath := filepath.Join(WorkspaceDir, "conf.json")
	if common.File.IsExist(confPath) {
		if data, err := os.ReadFile(confPath); nil != err {
			common.Log.Error("load conf [%s] failed: %s", confPath, err)
		} else {
			if err = json.Unmarshal(data, Conf); err != nil {
				common.Log.Error("parse conf [%s] failed: %s", confPath, err)
			} else {
				common.Log.Info("loaded conf [%s]", confPath)
			}
		}
	}
	Conf.setDefaultValue()
	common.Log.SetLogLevel(Conf.LogLevel)
	Conf.Save()
}

func (conf *AppConf) Save() {
	if ReadOnly {
		return
	}

	Conf.m.Lock()
	defer Conf.m.Unlock()

	newData, _ := json.MarshalIndent(Conf, "", "  ")
	oldData, err := common.FileLock.ReadFile(ConfigPath)
	if nil != err {
		conf.save(newData)
		return
	}

	if bytes.Equal(newData, oldData) {
		return
	}

	conf.save(newData)
}

func (conf *AppConf) save(data []byte) {
	if err := common.FileLock.WriteFile(ConfigPath, data); nil != err {
		common.Log.Error("write conf [%s] failed: %s", ConfigPath, err)
		return
	}
}

func (conf *AppConf) Close() {
	conf.Save()
}

type NetworkProxy struct {
	Scheme string `json:"scheme"`
	Host   string `json:"host"`
	Port   string `json:"port"`
}

func (np *NetworkProxy) String() string {
	if "" == np.Scheme {
		return ""
	}
	return np.Scheme + "://" + np.Host + ":" + np.Port
}

func (conf *AppConf) setDefaultValue() {
	if conf.SSL {
		if conf.CertFileName == "" {
			conf.CertFileName = "cert.pem"
		}
		if conf.CertKeyFileName == "" {
			conf.CertKeyFileName = "key.pem"
		}
	}
}
