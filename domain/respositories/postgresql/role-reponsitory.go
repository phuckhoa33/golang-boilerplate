package respositories

import (
	"github.com/jinzhu/gorm"
	"golang-boilerplate/domain/models/postgresql"
	"golang-boilerplate/domain/respositories/abstracts/base"
)

type IRoleRepository interface {
	*base.IBaseRepository
	GetRoleByName(role *postgresql.Role, name string)
}

type RoleRepository struct {
	*gorm.DB
	*base.RepositoryAbstract
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		DB:                 db,
		RepositoryAbstract: base.NewBaseRepositoryAbstract(db),
	}
}
