package mail

import (
	"errors"
	"net/smtp"
	"net/url")


type smtpProvider struct {
	Mail
	plainAuth smtp.Auth
}


func CreateSMTP(url *url.URL) (*smtpProvider, error) {
	username := url.User.Username()
	password, isSet := url.User.Password()

	if false == isSet {
		return nil,	errors.New("password not set")
	}
	auth := smtp.PlainAuth("", username, password, url.Host)

	return &smtpProvider{plainAuth: auth}, nil
}

func (s *smtpProvider) Send(email, title, userID string, body []byte) error {
	
	return nil
}