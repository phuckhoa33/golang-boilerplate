package config

import (
	"os"
	"strconv"
)

type MailConfig struct {
	MailHost     string
	MailPort     int
	MailUsername string
	MailPassword string
}

func LoadMailConfig() MailConfig {

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic(err)
	}

	return MailConfig{
		MailHost:     os.Getenv("MAIL_HOST"),
		MailPort:     port,
		MailUsername: os.Getenv("MAIL_USERNAME"),
		MailPassword: os.Getenv("MAIL_PASSWORD"),
	}
}
