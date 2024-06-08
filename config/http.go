package config

import "os"

type HTTPConfig struct {
	AppPort string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		AppPort: os.Getenv("APP_PORT"),
	}
}
