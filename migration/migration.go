package main

import (
	"fmt"
	"golang-boilerplate/config"
	db "golang-boilerplate/db/postgres"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq" // Postgres driver
)

type Migration struct {
}

func main() {
	config := config.NewConfig()
	database := db.Init(config)

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// Specify the directory containing your SQL migration files
	migrationsDir := filepath.Join(wd, "migration", "postgresql")
	fmt.Println(os.Getwd())
	// Read the directory contents
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Failed to list files in directory: %v", err)
	}

	// Iterate over each file in the directory
	for _, file := range files {
		// Check if the file is a SQL file
		if filepath.Ext(file.Name()) == ".sql" {
			fmt.Printf("Executing migration file: %s\n", file.Name())

			// Construct the full path to the file
			filePath := filepath.Join(migrationsDir, file.Name())

			// Read the SQL file
			sqlContent, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Fatalf("Failed to read migration file: %v", err)
			}

			// Execute the SQL commands
			_, err = database.DB().Exec(string(sqlContent))
			if err != nil {
				log.Fatalf("Failed to execute migration: %v", err)
			}

			fmt.Println("Migration executed successfully")
		}
	}
}
