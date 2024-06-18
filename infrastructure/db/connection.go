package db

import (
	"fmt"
	"golang-boilerplate/config"
	"log"

	"github.com/jinzhu/gorm"
)

func Init(config *config.Config) *gorm.DB {

	// Assign database source string
	databaseSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.Db.DatabaseHost,
		config.Db.DatabasePort,
		config.Db.DatabaseName,
		config.Db.DatabaseUser,
		config.Db.DatabasePassword)
	log.Printf("Database run as %s", databaseSourceName)

	// Connnect database with gorm
	db, err := gorm.Open("postgres", databaseSourceName)
	// Check error
	if err != nil {
		panic(err.Error())
	}

	return db
}
