package respositories

import (
	"github.com/jinzhu/gorm"
	"golang-boilerplate/domain/respositories/abstracts/base"
)

// IUserRepository is a contract what this repository can do
type IUserRepository interface {
	base.IBaseRepository
}

// UserRepository is a struct that represents the repository for the user model
type UserRepository struct {
	*gorm.DB
	*base.RepositoryAbstract
}

// NewUserRepository is a function that returns a new instance of the UserRepository struct
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB:                 db,
		RepositoryAbstract: base.NewBaseRepositoryAbstract(db),
	}
}
