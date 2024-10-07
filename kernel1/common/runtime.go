package common

import (
	"github.com/denisbrodbeck/machineid"
	"os"
)

func (g *GuRuntime) GetDeviceID() string {
	machineID, err := machineid.ID()
	if nil != err {
		return ""
	}
	return machineID
}

func (g *GuRuntime) GetDeviceName() string {
	ret, err := os.Hostname()
	if nil != err {
		return "unknown"
	}
	return ret
}
