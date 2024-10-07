package model

import "gorm.io/gorm"

type Config struct {
	Model
	gorm.Model
	Env      string
	Device   string
	Type     string
	OnlyRead int
	Key      string
	Value    string
}
