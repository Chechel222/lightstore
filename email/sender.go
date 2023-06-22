package email

import (
	"bytes"
	"errors"
	"text/template"
)

type SendEmailInput struct {
	To      string
	Body    string
	Subject string
}

type Sender interface {
	Send(input SendEmailInput) error
}

func (e *SendEmailInput) GenerateBodyFromHTML(templateFileName string, data interface{}) error {

	tmpl, err := template.ParseFiles(templateFileName)

	if err != nil {

		return err
	}

	buffer := new(bytes.Buffer)
	if err = tmpl.Execute(buffer, data); err != nil {
		return err
	}

	e.Body = buffer.String()
	return nil
}

func (e *SendEmailInput) Validate() error {

	if e.To == "" {
		return errors.New("empty to")
	}

	if !IsEmailValid(e.To) {
		return errors.New("invalid to email")
	}

	return nil
}
