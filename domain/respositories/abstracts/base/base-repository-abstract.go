package base

import (
	"github.com/jinzhu/gorm"
)

/*
IBaseRepository is an interface that defines the basic methods that a repository should have.

Usage:
	- Create a new interface that extends IBaseRepository and add the methods that you need.
	- Create a struct that implements the interface and the IBaseRepository.
	- Implement the methods that you added in the interface.
	- Create a function that returns the struct that you created.
	- Use the methods that you implemented in the struct.
*/

// IBaseRepository is an interface that defines the basic methods that a repository should have.
type IBaseRepository interface {
	FindById(model interface{}, id interface{})
	Insert(model interface{})
	UpdateModel(model interface{}, request interface{})
	UpdateOne(model interface{}, property interface{}, value interface{})
	Delete(id interface{})
	FindAll(model interface{})
	FindOne(model interface{}, query interface{}, value interface{})
}

// RepositoryAbstract is a struct that implements the IBaseRepository interface.
type RepositoryAbstract struct {
	*gorm.DB
	IBaseRepository
}

// FindById retrieves a model by its ID from the database.
// It expects a model instance (to populate with the found record) and an ID value.
// The function uses the ID to locate the record in the database and populates the provided model instance.
func (bra *RepositoryAbstract) FindById(model interface{}, id interface{}) {
	bra.DB.Where("id = ?", id).Find(model)
}

// Insert adds a new record to the database based on the model provided.
// The model should be a struct corresponding to the table structure, populated with the values to insert.
func (bra *RepositoryAbstract) Insert(model interface{}) {
	db := bra.DB.Create(&model)
	if db.Error != nil {
		panic(db.Error)
	}
}

// UpdateModel updates fields of a model in the database.
// It accepts a model instance (which should have its primary key field populated) and a request object containing the fields to update.
func (bra *RepositoryAbstract) UpdateModel(model interface{}, request interface{}) {
	db := bra.DB.Model(&model).Update(request)
	if db.Error != nil {
		panic(db.Error)
	}
}

// UpdateOne updates a single property of a model in the database.
// It requires the model instance, the property name to update, and the new value for that property.
func (bra *RepositoryAbstract) UpdateOne(model interface{}, property interface{}, value interface{}) {
	db := bra.DB.Model(&model).Where(property, value).Update(property, value)
	if db.Error != nil {
		panic(db.Error)
	}
}

// Delete removes a record from the database by its ID.
// The function expects the ID of the record to delete.
func (bra *RepositoryAbstract) Delete(id interface{}) {
	bra.DB.Where("id = ?", id).Delete(id)
}

// FindAll retrieves all records for a given model from the database.
// It expects a model instance which will be populated with the results.
func (bra *RepositoryAbstract) FindAll(model interface{}) {
	bra.DB.Find(model)
}

// FindOne retrieves a single record from the database based on a query and a value.
// It expects a model instance which will be populated with the result.
func (bra *RepositoryAbstract) FindOne(model interface{}, query interface{}, value interface{}) {
	bra.DB.Where(query, value).Find(model)
}

// NewBaseRepositoryAbstract creates a new RepositoryAbstract instance with the provided database connection.
func NewBaseRepositoryAbstract(db *gorm.DB) *RepositoryAbstract {
	return &RepositoryAbstract{DB: db}
}
