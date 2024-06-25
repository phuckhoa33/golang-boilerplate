package models

type Role struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(200);"`
	Permision string `json:"permision" gorm:"type:text[];"`
	CreatedAt string `json:"createdAt" gorm:"type:varchar(200);"`
	UpdatedAt string `json:"updatedAt" gorm:"type:varchar(200);"`
	DeleteAt  string `json:"deleteAt" gorm:"type:varchar(200);"`
}
