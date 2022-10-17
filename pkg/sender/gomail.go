package sender

import (
	"fmt"
	"log"
	"net/smtp"
)

type GoMail struct {
	auth    smtp.Auth
	address string
	from    string
}
type Sender interface {
	Send(subject, body string, recipients ...string) error
}

func NewGoMail(host, port, from, password string) *GoMail {
	return &GoMail{
		address: host + ":" + port,
		auth:    smtp.PlainAuth("", from, password, host),
		from:    from,
	}
}

func (m *GoMail) Send(subject, body string, recipients ...string) error {
	err := smtp.SendMail(m.address,
		m.auth,
		m.from,
		recipients,
		m.getMessage(subject, body))

	if err != nil {
		log.Printf("smtp error: %s", err)

		return err
	}

	return nil
}

func (m *GoMail) getMessage(subject, body string) []byte {

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", m.from)
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += fmt.Sprintf("\r\n%s\r\n", body)

	return []byte(msg)
}
