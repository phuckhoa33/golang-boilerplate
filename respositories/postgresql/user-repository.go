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

func (userRepository *UserRepository) Create(user *models.User) {
	db := userRepository.DB.Create(&user)
	if db.Error != nil {
		panic(db.Error)
	}
}

func (userRepository *UserRepository) GetUserById(user *models.User, id any) {
	userRepository.DB.Where("id = ?", id).Find(user)
}

func (userRepository *UserRepository) UpdateSingleProperty(user *models.User, propertyName string, value string) {
	// Update value in users database depend on propertyName
	userRepository.DB.Model(&user).Update(propertyName, value)
}
