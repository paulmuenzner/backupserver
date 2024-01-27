package email

import "gopkg.in/gomail.v2"

/////////////////////////////////////////////////////////////////////////
// Setup of Dependency Injection for Transactional Email Client Methods
///////////////////////
type EmailMethods interface {
	SendEmail(senderEmail, recipientEmail, subject, body string) error
	SendEmailBackupSuccess(timeStampString)
}

type MailClients struct {
	MyEmailClient *gomail.Dialer
}

type MethodConfig struct {
	MethodInterface EmailMethods
}

func NewClientBasics(emailClient *MailClients) *MethodConfig {
	return &MethodConfig{MethodInterface: emailClient}
}
