package models

import "github.com/jinzhu/gorm"

type UserDevice struct {
	*gorm.Model
	ID           uint   `json:"id" gorm:"primary_key"`
	UserID       uint   `json:"user_id" gorm:"type:int;"`
	DeviceID     uint   `json:"device_id" gorm:"type:int;"`
	AccessType   string `json:"access_type" gorm:"type:varchar(200);"`
	Browser      string `json:"browser" gorm:"type:varchar(200);"`
	Os           string `json:"os" gorm:"type:varchar(200);"`
	DeviceStatus string `json:"device_status" gorm:"type:varchar(200);"`
	RefreshToken string `json:"refresh_token" gorm:"type:varchar(200);"`
	CreatedAt    string `json:"created_at" gorm:"type:varchar(200);"`
	UpdatedAt    string `json:"updated_at" gorm:"type:varchar(200);"`
	DeleteAt     string `json:"delete_at" gorm:"type:varchar(200);"`
}
