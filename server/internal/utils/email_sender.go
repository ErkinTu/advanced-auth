package utils

import (
	"AdvAuthGo/config"
	"fmt"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	From     string
	Password string
	Host     string
	Port     string
}

func NewEmailSender(cfg *config.Config) *EmailSender {
	return &EmailSender{
		From:     cfg.SMTPEmail,
		Password: cfg.SMTPPass,
		Host:     cfg.SMTPHost,
		Port:     cfg.SMTPPort,
	}
}

func (s *EmailSender) SendActivationEmail(to, activationURL string) error {
	port, err := strconv.Atoi(s.Port)
	if err != nil {
		return fmt.Errorf("invalid port value: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Activate your account")

	body := "Click to activate your account:\n" + activationURL
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(s.Host, port, s.From, s.Password)

	return d.DialAndSend(m)
}
