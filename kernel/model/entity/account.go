package entity

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Title    string `json:"title"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Remark   string `json:"remark"`
	Type     string `json:"type"`
	Group    string `json:"group"`
}
