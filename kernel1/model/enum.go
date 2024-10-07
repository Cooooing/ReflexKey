package model

import "kernel/common"

type enum interface {
	String() string
	Index() int
}

type Enum int

const (
	ConfigTypeEnum = iota
)

var ConfigTypeEnumStr = [][]string{ConfigTypeStr}

func (e Enum) String() []string {
	return ConfigTypeEnumStr[e]
}

func (e Enum) Index() int {
	return int(e)
}

// EnumValues 返回枚举的所有值
func EnumValues() [][]string {
	return ConfigTypeEnumStr
}

// Values 返回枚举的所有值
func Values(e Enum) []string {
	return e.String()
}

// ExistOf 判断某值是否存在枚举值中
func ExistOf(e enum, str string) bool {
	common.Log.Info(e.String())
	for _, v := range e.String() {
		common.Log.Info("v: %d", v)
		//if v == str {
		//	return true
		//}
	}
	return false
}

// ConfigType 配置类型
type ConfigType int

// 声明每个枚举项的索引值
const (
	Base ConfigType = iota
	User
)

var ConfigTypeStr = []string{"Base", "User"}

// String 返回枚举项的索引值
func (w ConfigType) String() string {
	return ConfigTypeStr[w]
}

// Index 返回枚举项的字符值
func (w ConfigType) Index() int {
	return int(w)
}
