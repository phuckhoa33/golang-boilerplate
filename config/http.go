package config

import "os"

type HTTPConfig struct {
	SwaggerEnable  string
	AppName        string
	AppDescription string
	AppVersion     string
	AppEnv         string
	AppApiDebug    string
	AppApiPrefix   string
	AppPort        string
	AppHost        string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		SwaggerEnable:  os.Getenv("SWAGGER_ENABLE"),
		AppName:        os.Getenv("APP_NAME"),
		AppDescription: os.Getenv("APP_DESCRIPTION"),
		AppVersion:     os.Getenv("APP_VERSION"),
		AppEnv:         os.Getenv("APP_ENV"),
		AppApiDebug:    os.Getenv("APP_API_DEBUG"),
		AppApiPrefix:   os.Getenv("APP_API_PREFIX"),
		AppPort:        os.Getenv("APP_PORT"),
		AppHost:        os.Getenv("APP_HOST"),
	}
}
