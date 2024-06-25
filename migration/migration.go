package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/pressly/goose"
)

type Migration struct {
}

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	err = goose.SetDialect("postgres")
	if err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	err = goose.Up(db, "postgres")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
