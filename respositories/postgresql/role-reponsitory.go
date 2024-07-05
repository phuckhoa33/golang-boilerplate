package respositories

import (
	"github.com/jinzhu/gorm"
	"golang-boilerplate/models"
)

type RoleRepositoryQ interface {
	GetRoleByName(role *models.Role, name string)
}

type RoleRepository struct {
	*gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (roleRepository *RoleRepository) GetRoleByName(role *models.Role, name string) {
	roleRepository.DB.Where("name = ?", name).Find(role)
}
