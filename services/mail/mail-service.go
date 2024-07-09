package mail_service

import (
	"bytes"
	"golang-boilerplate/config"
	"html/template"

	"gopkg.in/gomail.v2"
)

type MailService struct {
	config *config.Config
}

func NewMailService(config *config.Config) *MailService {
	return &MailService{
		config: config,
	}
}

func (ms *MailService) SendEmail(to, subject, templateFile string, data interface{}) error {
	// Parse the email template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	// Render the email template with the provided data
	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, data)
	if err != nil {
		return err
	}

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", ms.config.Mail.MailUsername)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", bodyContent.String())

	// Create a new SMTP client
	d := gomail.NewDialer(ms.config.Mail.MailHost, ms.config.Mail.MailPort, ms.config.Mail.MailUsername, ms.config.Mail.MailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
