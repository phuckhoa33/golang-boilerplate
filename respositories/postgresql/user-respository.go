package respositories

import (
	"golang-boilerplate/models"

	"github.com/jinzhu/gorm"
)

type UserRepositoryQ interface {
	GetUserByEmail(user *models.User, email string)
}

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepository *UserRepository) GetUserByEmail(user *models.User, email string) {
	userRepository.DB.Where("email = ?", email).Find(user)
}
