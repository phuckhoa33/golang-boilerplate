package services

import (
	"golang-boilerplate/config"
	mail_service "golang-boilerplate/services/mail"
)

type ServicesInitialization struct {
	MailServiceConfig *mail_service.MailService
}

func NewServicesInitialization(config *config.Config) ServicesInitialization {
	return ServicesInitialization{
		MailServiceConfig: mail_service.NewMailService(config),
	}
}
