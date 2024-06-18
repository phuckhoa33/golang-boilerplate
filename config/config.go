package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	App  AppConfig
	Db   DbConfig
	Auth AuthConfig
}

func NewConfig() *Config {
	// Load environment variables and check error
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		App:  LoadAppConfig(),
		Db:   LoadDbConfig(),
		Auth: LoadAuthConfig(),
	}
}
