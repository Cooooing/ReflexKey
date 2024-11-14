package sql

import (
	"kernel/model"
)

const key = "reflex_key_database_ver"

func getDatabaseVer() (ret string) {
	var config model.Config
	DB.Where("env = ? and device = ? and key = ?", "all", "all", key).First(&config)
	return config.Value
}
