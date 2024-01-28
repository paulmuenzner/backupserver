package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func NewEmailClient(configData *EmailClientConfigData) (client *MailClient, err error) {
	dialer := gomail.NewDialer(configData.Host, configData.SmtpPort, configData.SmtpUsername, configData.SmtpPassword)
	if dialer == nil {
		return nil, fmt.Errorf("Failed to create dialer in 'NewEmailClient()' using email client config data: %+v", configData)
	}

	return &MailClient{MyEmailClient: dialer}, nil
}
