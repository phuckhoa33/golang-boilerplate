package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	App   AppConfig
	DB    DBConfig
	Auth  AuthConfig
	Mail  MailConfig
	Minio MinioConfig
}

func NewConfig() *Config {
	// Load environment variables and check error
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	env := &Config{
		App:   LoadAppConfig(),
		DB:    LoadDBConfig(),
		Auth:  LoadAuthConfig(),
		Mail:  LoadMailConfig(),
		Minio: LoadMinioConfig(),
	}

	if env.App.AppEnv == "development" {
		log.Println("The App is running in development app")
	}

	return env
}
