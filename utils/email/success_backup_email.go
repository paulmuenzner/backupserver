package email

import (
	"fmt"
	"time"

	"github.com/paulmuenzner/golang-backupserver/config"
	"github.com/paulmuenzner/golang-backupserver/utils/date"
	envHandler "github.com/paulmuenzner/golang-backupserver/utils/env"
	strings "github.com/paulmuenzner/golang-backupserver/utils/strings"
)

func (client *MailClient) SendEmailBackupSuccess(timeStamp time.Time, bucketName, folderPathBackup, databaseName string) error {
	timeStampString := date.TimeStampSlug(timeStamp)
	// Subject
	subjectComponents := []string{"Successful backup: ", timeStampString}
	subject := strings.ConcatenateStrings(subjectComponents...)

	// Subject
	bodyComponents := []string{"Backup of your database successful. <br/> Date: ", timeStampString, "<br/> Database name: ", databaseName, "<br/> Bucket name: ", bucketName, "<br/> Folder path S3: ", folderPathBackup}
	body := strings.ConcatenateStrings(bodyComponents...)

	senderEmail, err := envHandler.GetEnvValue(config.EmailAddressSenderEnv, "")
	recipientEmail, err := envHandler.GetEnvValue(config.EmailAddressReceiverEnv, "")
	// Log as error if no defaultValue provided in GetEnvValue()
	if err != nil {
		return fmt.Errorf("Error in 'SendEmailBackupSuccess' applying 'GetEnvValue()'. Cannot retrieve env value. Error: %v", err)
	}

	// Send
	err = client.SendEmail(senderEmail, recipientEmail, subject, body)
	if err != nil {
		return fmt.Errorf("Error in 'SendEmailBackupSuccess' applying 'SendEmail()'. Client data: %+v. Error: %v", client, err)
	}
	return nil
}
