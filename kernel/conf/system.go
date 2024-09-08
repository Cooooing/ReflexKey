package conf

type System struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	KernelVersion    string `json:"kernelVersion"`
	OS               string `json:"os"`
	OSPlatform       string `json:"osPlatform"`
	Container        string `json:"container"` // docker, android, ios, std
	IsMicrosoftStore bool   `json:"isMicrosoftStore"`
	IsInsider        bool   `json:"isInsider"`

	HomeDir      string `json:"homeDir"`
	WorkspaceDir string `json:"workspaceDir"`
	AppDir       string `json:"appDir"`
	ConfDir      string `json:"confDir"`
	DataDir      string `json:"dataDir"`

	NetworkServe bool          `json:"networkServe"` // 是否开启网络伺服
	NetworkProxy *NetworkProxy `json:"networkProxy"`

	UploadErrLog           bool `json:"uploadErrLog"`
	DisableGoogleAnalytics bool `json:"disableGoogleAnalytics"`
	DownloadInstallPkg     bool `json:"downloadInstallPkg"`
	AutoLaunch2            int  `json:"autoLaunch2"`    // 0：不自动启动，1：自动启动，2：自动启动+隐藏主窗口
	LockScreenMode         int  `json:"lockScreenMode"` // 0：手动，1：手动+跟随系统
}

func NewSystem() *System {
	return &System{
		//ID:                 util.GetDeviceID(),
		//Name:               util.GetDeviceName(),
		KernelVersion:      VERSION,
		NetworkProxy:       &NetworkProxy{},
		DownloadInstallPkg: true,
	}
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
