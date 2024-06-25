package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email       string    `json:"email" gorm:"type:varchar(200);"`
	Username    string    `json:"username" gorm:"type:varchar(200);"`
	PhoneNumber string    `json:"phoneNumber" gorm:"type:varchar(200);"`
	Fullname    string    `json:"fullname" gorm:"type:varchar(200);"`
	Address     string    `json:"address" gorm:"type:varchar(200);"`
	Gender      string    `json:"gender" gorm:"type:varchar(200);"`
	DateOfBirth time.Time `json:"dateOfBirth" gorm:"type:date;"`
	Password    string    `json:"password" gorm:"type:varchar(200);"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(200);"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:date;"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:date;"`
}
