package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId   string `json:"userId" gorm: "type:varchar(200);"`
	Email    string `json:"email" gorm:"type:varchar(200);"`
	Name     string `json:"name" gorm:"type:varchar(200);"`
	Password string `json:"password" gorm:"type:varchar(200);"`
}
