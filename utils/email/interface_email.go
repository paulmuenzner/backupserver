package email

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

// ///////////////////////////////////////////////////////////////////////
// Setup of Dependency Injection for Transactional Email Client Methods
// /////////////////////
type EmailMethods interface {
	SendEmailBackupSuccess(timeStamp time.Time, bucketName, folderPathBackup, databaseName string) error
	SendEmail(senderEmail, recipientEmail, subject, body string) error
}

type MailClient struct {
	MyEmailClient *gomail.Dialer
}

type EmailClientConfigData struct {
	Host         string
	SmtpUsername string
	SmtpPassword string
	SmtpPort     int
}

type EmailMethodInterface struct {
	MethodInterface EmailMethods
}

func NewEmailMetodInterface(emailClient *MailClient) *EmailMethodInterface {
	return &EmailMethodInterface{MethodInterface: emailClient}
}

func GetEmailMethods(emailClientConfig *EmailClientConfigData) (emailClientMethods *EmailMethodInterface, err error) {
	// Setup email client dependency
	client, err := NewEmailClient(emailClientConfig)
	if err != nil {
		return nil, fmt.Errorf("Cannot create email client in 'EmailProductionClient()' with 'NewEmailClient()'. Email client config: %+v. Error: %v", emailClientConfig, err)
	}
	emailClientMethods = NewEmailMetodInterface(client)

	return emailClientMethods, nil
}
