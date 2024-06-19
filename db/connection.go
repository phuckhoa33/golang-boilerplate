package db

import (
	"fmt"
	"golang-boilerplate/config"
	"log"

	"github.com/jinzhu/gorm"
)

func Init(config *config.Config) *gorm.DB {
	databaseSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.DB.PostgresConfig.PostgresDatabaseHost,
		config.DB.PostgresConfig.PostgresDatabasePort,
		config.DB.PostgresConfig.PostgresDatabaseName,
		config.DB.PostgresConfig.PostgresDatabaseUser,
		config.DB.PostgresConfig.PostgresDatabasePassword)

	log.Printf("Database run as %s", databaseSourceName)

	// Connnect database with gorm
	db, err := gorm.Open("postgres", databaseSourceName)
	// Check error
	if err != nil {
		panic(err.Error())
	}

	return db
}
