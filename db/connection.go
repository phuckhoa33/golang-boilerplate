package db

import (
	"fmt"
	"golang-boilerplate/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(config *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		config.DB.DatabaseUser,
		config.DB.DatabasePassword,
		config.DB.DatabaseHost,
		config.DB.DatabaseName)

	fmt.Println(dataSourceName)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
