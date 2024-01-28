package email

import "gopkg.in/gomail.v2"

func (client *MailClient) SendEmail(senderEmail, recipientEmail, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Send
	if err := client.MyEmailClient.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
