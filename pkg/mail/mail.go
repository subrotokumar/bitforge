package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAdddress  = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, from, password string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  from,
		fromEmailPassword: password,
	}
}

func (gmailSender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", gmailSender.name, gmailSender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("Error in attracting file %s: %w", f, err)
		}
	}
	smtpAuth := smtp.PlainAuth("", gmailSender.fromEmailAddress, gmailSender.fromEmailPassword, smtpAuthAdddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
