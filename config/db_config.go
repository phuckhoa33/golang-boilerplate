package config

import "os"

type DbConfig struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseDebug    string
}

func LoadDbConfig() DbConfig {
	return DbConfig{
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseDebug:    os.Getenv("DATABASE_DEBUG"),
	}
}
