package config

import "os"

type AppConfig struct {
	SwaggerEnable  string
	AppName        string
	AppDescription string
	AppVersion     string
	AppEnv         string
	AppPort        string
	AppApiDebug    string
	AppApiPrefix   string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		SwaggerEnable:  os.Getenv("SWAGGER_ENABLE"),
		AppName:        os.Getenv("APP_NAME"),
		AppDescription: os.Getenv("APP_DESCRIPTION"),
		AppVersion:     os.Getenv("APP_VERSION"),
		AppEnv:         os.Getenv("APP_ENV"),
		AppPort:        os.Getenv("APP_PORT"),
		AppApiDebug:    os.Getenv("APP_API_DEBUG"),
		AppApiPrefix:   os.Getenv("APP_API_PREFIX"),
	}
}
