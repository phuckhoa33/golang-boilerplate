package postgresql

import (
	"golang-boilerplate/domain/models/abstracts/base"
)

type UserDevice struct {
	*base.AuditModelAbstract
	UserID       uint   `json:"userId" gorm:"type:int;"`
	DeviceID     uint   `json:"deviceId" gorm:"type:int;"`
	AccessType   string `json:"accessType" gorm:"type:varchar(200);"`
	Browser      string `json:"browser" gorm:"type:varchar(200);"`
	Os           string `json:"os" gorm:"type:varchar(200);"`
	DeviceStatus string `json:"deviceStatus" gorm:"type:varchar(200);"`
	RefreshToken string `json:"refreshToken" gorm:"type:varchar(200);"`
}
