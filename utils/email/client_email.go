package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func NewEmailClient(emailClientConfig *EmailClientConfigData) (client *MailClient, err error) {
	dialer := gomail.NewDialer(emailClientConfig.Host, emailClientConfig.SmtpPort, emailClientConfig.SmtpUsername, emailClientConfig.SmtpPassword)
	if dialer == nil {
		return nil, fmt.Errorf("Failed to create dialer in 'NewEmailClient()' using email client config data.")
	}

	return &MailClient{MyEmailClient: dialer}, nil
}
