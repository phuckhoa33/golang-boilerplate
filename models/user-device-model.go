package models

import "github.com/google/uuid"

type UserDevice struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	UserID       uint      `json:"userId" gorm:"type:int;"`
	DeviceID     uint      `json:"deviceId" gorm:"type:int;"`
	AccessType   string    `json:"accessType" gorm:"type:varchar(200);"`
	Browser      string    `json:"browser" gorm:"type:varchar(200);"`
	Os           string    `json:"os" gorm:"type:varchar(200);"`
	DeviceStatus string    `json:"deviceStatus" gorm:"type:varchar(200);"`
	RefreshToken string    `json:"refreshToken" gorm:"type:varchar(200);"`
	CreatedAt    string    `json:"createdAt" gorm:"type:varchar(200);"`
	UpdatedAt    string    `json:"updatedAt" gorm:"type:varchar(200);"`
	DeleteAt     string    `json:"deleteAt" gorm:"type:varchar(200);"`
}
