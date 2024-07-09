package postgresql

import (
	"golang-boilerplate/domain/models/abstracts/base"
)

type Role struct {
	*base.FullAuditModelAbstract
	Name        string `json:"name" gorm:"type:varchar(200);"`
	Permissions string `json:"permissions" gorm:"type:text[];"`
}
