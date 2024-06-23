package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email     string    `json:"email" gorm:"type:varchar(200);"`
	Name      string    `json:"name" gorm:"type:varchar(200);"`
	Password  string    `json:"password" gorm:"type:varchar(200);"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(200);"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:date;"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:date;"`
}
