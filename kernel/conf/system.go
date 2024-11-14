package conf

import (
	"kernel/common"
)

type System struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	OS         string `json:"os"`
	OSPlatform string `json:"osPlatform"`
	Container  string `json:"container"` // docker, android, ios, std

}

func NewSystem() *System {
	return &System{
		ID:   common.Runtime.GetDeviceID(),
		Name: common.Runtime.GetDeviceName(),
	}
}
