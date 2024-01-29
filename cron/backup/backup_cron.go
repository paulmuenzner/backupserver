package backup

import (
	"github.com/paulmuenzner/golang-backupserver/config"
	services "github.com/paulmuenzner/golang-backupserver/services"
	"github.com/paulmuenzner/golang-backupserver/utils/aws"
	date "github.com/paulmuenzner/golang-backupserver/utils/date"
	email "github.com/paulmuenzner/golang-backupserver/utils/email"
	files "github.com/paulmuenzner/golang-backupserver/utils/files"
	logger "github.com/paulmuenzner/golang-backupserver/utils/logs"
	mongoDB "github.com/paulmuenzner/golang-backupserver/utils/mongoDB"

	"os"
)

// More on cron jobs in Go -> https://pkg.go.dev/github.com/robfig/cron#section-readme
func Backup(mongoDbClientConfig *mongoDB.DatabaseClient, awsClientConfig *aws.AwsClientConfigData, emailClientConfig *email.EmailClientConfigData, bucketName string) {
	//////////////////////////////////////
	// Define time stamps
	timeStamp := date.TimeStamp()
	timeStampString := date.TimeStampSlug(timeStamp)

	// Generate a distinct backup folder named with the current timestamp
	folderPathBackup := files.GetSubFolder(timeStampString)
	err := os.MkdirAll(folderPathBackup, os.ModePerm)
	if err != nil {
		logger.GetLogger().Error("Error creating subfolder in 'Backup()'.", " Folder path: ", folderPathBackup, ". Error: ", err)
		return
	}

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Generate a CSV file containing metadata details, such as timestamp, size, and full file path, for all created backup files
	fileNameMeta := config.FileNameMetaData
	err = services.CreateMetaDataFile(folderPathBackup, fileNameMeta)
	if err != nil {
		logger.GetLogger().Error("Error creating meta data file in 'Backup()'.", " Meta file name: ", fileNameMeta, " Folder path: ", folderPathBackup, ". Error: ", err)
		return
	}

	/////////////////////////////////////////////////////////////////
	// Setup database client for following Dependency Injections
	databaseName := config.NameDatabase
	databaseClientSetup := mongoDB.MongoClientBasics(mongoDbClientConfig)

	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Loop through all database collections and create backup files in local backup folder 'folderPathBackup'
	err = CreateBackupFiles(databaseClientSetup, databaseName, timeStamp, folderPathBackup, fileNameMeta)
	if err != nil {
		logger.GetLogger().Error("Error creating all backup files in 'Backup()' using 'CreateBackupFiles()'. Error: ", err)
		return
	}

	/////////////////////////////////////////////////////////////////////////////////////////
	// Manage Storage Backup
	////////////////////////
	err = ManageStorages(folderPathBackup, fileNameMeta, awsClientConfig, bucketName)
	if err != nil {
		logger.GetLogger().Error("Error in 'Backup()' applying 'ManageStorages()'. Error: ", err)
		return
	}

	if config.SendEmailNotifications == true {
		// Setup Email client dependency
		emailMethods, err := email.GetEmailMethods(emailClientConfig)
		err = emailMethods.MethodInterface.SendEmailBackupSuccess(timeStamp, bucketName, folderPathBackup, databaseName)
		if err != nil {
			logger.GetLogger().Error("Error in 'Backup()' applying 'SendEmailBackupSuccess()'. Error: ", err)
			return
		}
	}

	// Log successful backup
	logger.GetLogger().Infof("Backup successful. Date: %s. Database name: %s. S3 folder path: %s. Bucket name: %s. Meta file name: %s", timeStampString, databaseName, folderPathBackup, bucketName, fileNameMeta)
}
