package model

import (
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Env      string
	Device   string
	Type     string
	OnlyRead int
	Key      string
	Value    string
}
