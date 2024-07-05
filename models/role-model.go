package models

import "github.com/google/uuid"

type Role struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name        string    `json:"name" gorm:"type:varchar(200);"`
	Permissions string    `json:"permissions" gorm:"type:text[];"`
	CreatedAt   string    `json:"createdAt" gorm:"type:varchar(200);"`
	UpdatedAt   string    `json:"updatedAt" gorm:"type:varchar(200);"`
	DeleteAt    string    `json:"deleteAt" gorm:"type:varchar(200);"`
}
