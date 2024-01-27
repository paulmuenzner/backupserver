package email

func (client *MailClients) SendEmailBackupSuccess(senderEmail, recipientEmail, subject, body string) error {
	err := client.SendEmail(senderEmail, recipientEmail, subject, body)

	// Send
	if err != nil {
		return err
	}
	return nil
}
