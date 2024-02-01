package email

import (
	"fmt"
	"time"

	"github.com/paulmuenzner/backupserver/config"
	"github.com/paulmuenzner/backupserver/utils/date"
	envHandler "github.com/paulmuenzner/backupserver/utils/env"
	strings "github.com/paulmuenzner/backupserver/utils/strings"
)

func (client *MailClient) SendEmailBackupSuccess(timeStamp time.Time, bucketName, folderPathBackup, databaseName string) error {
	timeStampString := date.TimeStampToUSFormat(timeStamp)
	// Subject
	subjectComponents := []string{"Successful backup: ", timeStampString}
	subject := strings.ConcatenateStrings(subjectComponents...)

	// Subject
	bodyComponents := []string{"<html><body><h1>Successful Database Backup.</h2> <br/><br/> Date: ", timeStampString, "<br/> Database name: ", databaseName, "<br/> Bucket name: ", bucketName, "<br/> Folder path S3: ", folderPathBackup, "</body></html>"}
	body := strings.ConcatenateStrings(bodyComponents...)

	senderEmailAddress, err := envHandler.GetEnvValue(config.EmailAddressSenderEnv, "") // Feel free to use default value via base_config
	// Log as error if no defaultValue provided in GetEnvValue()
	if err != nil {
		return fmt.Errorf("Error in 'SendEmailBackupSuccess()' utilizing 'GetEnvValue()' for 'senderEmailAddress'. Cannot retrieve env value. Error: %v", err)
	}

	recipientEmailAddress, err := envHandler.GetEnvValue(config.EmailAddressReceiverEnv, "") // Feel free to use default value via base_config
	// Log as error if no defaultValue provided in GetEnvValue()
	if err != nil {
		return fmt.Errorf("Error in 'SendEmailBackupSuccess()' utilizing 'GetEnvValue()' for 'recipientEmailAddress'. Cannot retrieve env value. Error: %v", err)
	}

	// Send
	err = client.SendEmail(senderEmailAddress, recipientEmailAddress, subject, body)
	if err != nil {
		return fmt.Errorf("Error in 'SendEmailBackupSuccess()' utilizing 'SendEmail()'. Client data: %+v. Error: %v", client, err)
	}
	return nil
}
