package respositories

import (
	"github.com/jinzhu/gorm"
	"golang-boilerplate/domain/models/postgresql"
	"golang-boilerplate/domain/respositories/abstracts/base"
)

type IUserRepository interface {
	GetUserByEmail(user *postgresql.User, email string)
	base.IBaseRepository
}

type UserRepository struct {
	*gorm.DB
	*base.RepositoryAbstract
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB:                 db,
		RepositoryAbstract: base.NewBaseRepositoryAbstract(db),
	}
}

func (userRepository *UserRepository) GetUserByEmail(user *postgresql.User, email string) {
	userRepository.DB.Where("email = ?", email).Find(user)
}
