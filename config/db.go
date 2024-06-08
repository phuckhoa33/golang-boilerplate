package config

import (
	"os"
)

type DBConfig struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseDebug    string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseDebug:    os.Getenv("DATABASE_DEBUG"),
	}
}
