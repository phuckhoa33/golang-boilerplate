package models

import (
	"github.com/google/uuid"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"` // This is the primary key
	Email            string    `json:"email" gorm:"type:varchar(200);"`
	Username         string    `json:"username" gorm:"type:varchar(200);"`
	PhoneNumber      string    `json:"phoneNumber" gorm:"type:varchar(200);"`
	FullName         string    `json:"fullName" gorm:"type:varchar(200);"`
	Address          string    `json:"address" gorm:"type:varchar(200);"`
	Gender           string    `json:"gender" gorm:"type:varchar(200);"`
	DateOfBirth      time.Time `json:"dateOfBirth" gorm:"type:date;"`
	Password         string    `json:"password" gorm:"type:varchar(200);"`
	Avatar           string    `json:"avatar" gorm:"type:varchar(200);"`
	RoleId           uuid.UUID `json:"roleId" gorm:"type:uuid;"`
	VerifyAccountOtp string    `json:"verifyAccountOtp" gorm:"type:varchar(6);"`
}
