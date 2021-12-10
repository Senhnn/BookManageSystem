package module

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	PowerLevel uint8  `json:"power_level"`
	RealName   string `json:"real_name"`
	Age        uint16 `json:"age"`
	Account    string `json:"account"`
	UserName   string `json:"user_name"`
}
