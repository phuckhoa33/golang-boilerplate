package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"` // This is the primary key
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
