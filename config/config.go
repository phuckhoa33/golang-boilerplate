package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	App  AppConfig
	DB   DBConfig
	Auth AuthConfig
}

func NewConfig() *Config {
	// Load environment variables and check error
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		App:  LoadAppConfig(),
		DB:   LoadDBConfig(),
		Auth: LoadAuthConfig(),
	}
}
