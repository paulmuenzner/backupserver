package email

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

// ///////////////////////////////////////////////////////////////////////
// Setup interface for email repository utilizing Dependency Injection
// /////////////////////
type EmailRepository interface {
	SendEmailBackupSuccess(timeStamp time.Time, bucketName, folderPathBackup, databaseName string) error
	SendEmailFailedBackup(timeStamp time.Time, errorMessage error, bucketName, folderPathBackup, databaseName string) error
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
	MethodInterface EmailRepository
}

func NewEmailMetodInterface(emailClient *MailClient) *EmailMethodInterface {
	return &EmailMethodInterface{MethodInterface: emailClient}
}

func GetEmailMethods(emailClientConfig *EmailClientConfigData) (emailClientMethods *EmailMethodInterface, err error) {
	// Setup email client dependency
	client, err := NewEmailClient(emailClientConfig)
	if err != nil {
		return nil, fmt.Errorf("Cannot create email client in 'EmailProductionClient()' with 'NewEmailClient()'. Error: %v", err)
	}
	emailClientMethods = NewEmailMetodInterface(client)

	return emailClientMethods, nil
}
