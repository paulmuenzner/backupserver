package cronBackup

import (
	"backupserver/config"
	services "backupserver/services"
	"backupserver/utils/aws"
	data "backupserver/utils/csv"
	date "backupserver/utils/date"
	email "backupserver/utils/email"
	files "backupserver/utils/files"
	logger "backupserver/utils/logs"
	mongoDB "backupserver/utils/mongoDB"

	"math"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

// More on cron jobs in Go -> https://pkg.go.dev/github.com/robfig/cron#section-readme
func Backup(mongoDbClientConfig *mongoDB.DatabaseClient, awsClientConfig *aws.AwsClientConfigData, emailClientConfig *email.EmailClientConfigData, bucketName string) {
	//////////////////////////////////////
	// Define time stamps
	timeStamp := date.TimeStamp()
	timeStampString := date.TimeStampSlug(timeStamp)

	// Create sub folder
	folderPathBackup := files.GetSubFolder(timeStampString)
	err := os.MkdirAll(folderPathBackup, os.ModePerm)
	if err != nil {
		logger.GetLogger().Error("Error creating subfolder in 'Backup'.", " Folder path: ", folderPathBackup, ". Error: ", err)
		return
	}

	// Create csv file for meta data
	fileNameMeta := config.FileNameMetaData
	err = services.CreateMetaDataFile(folderPathBackup, fileNameMeta)
	if err != nil {
		logger.GetLogger().Error("Error creating meta data file in 'Backup'.", " Meta file name: ", fileNameMeta, " Folder path: ", folderPathBackup, ". Error: ", err)
		return
	}

	//////////////////////////////////////
	// Get all collections in database
	databaseName := config.NameDatabase
	databaseClientSetup := mongoDB.MongoClientBasics(mongoDbClientConfig)
	collections, err := databaseClientSetup.MethodInterface.GetAllCollections(databaseName)
	if err != nil {
		logger.GetLogger().Error("Error adding in 'Backup' retrieving collection list. Error: ", err)
		return
	}

	//////////////////////////////////////
	// Loop through collections (slice)
	for _, collectionName := range collections {

		// Count number of documents in MongoDB collection
		totalDocuments, err := databaseClientSetup.MethodInterface.CountDocumentsInMongo(databaseName, collectionName)

		if err != nil {
			logger.GetLogger().Error("Error adding in 'Backup' counting documents in collection: ", collectionName, "Error: ", err)
			return
		}

		fileNameBackupFile := collectionName
		fileNumber := 0                                                           // Numerate backup files if more than one backup file is needed per collection due to permitted size limits of backup files
		interval := 100                                                           // number of documents retrieved per iteration step
		iterations := int(math.Ceil(float64(totalDocuments) / float64(interval))) // Number of iteration steps

		// Iterate through number of iteration steps to retrieve documents step-by-step inside MongoDB collection in moderate step size (interval)
		// Avoids to load too large data into memory in case of very huge MongoDB collections
		for i := 0; i < iterations; i++ {
			skip := i * interval

			// Retrieve limited number of documents and skip as needed
			var backupData []bson.M // prepare []interface{}
			err := databaseClientSetup.MethodInterface.FindDocumentsInMongo(databaseName, collectionName, interval, skip, &backupData)
			if err != nil {
				logger.GetLogger().Error("Error in 'Backup' retrieving documents in collection: ", collectionName, ". Skip: ", skip, ". Error: ", err)
				return
			}

			// Check if current backup file together with new data to add (backupData) already exists
			backupFileCreated, err := files.LocalFileExists(fileNameBackupFile)
			if err != nil {
				logger.GetLogger().Errorf("Error validating if file path exists in 'Backup' using 'LocalFileExists(fileNameBackupFile)'. File path %s. Error: %v", fileNameBackupFile, err)
				return
			}

			var backupFileSizeExceeded bool = false
			if backupFileCreated == true {
				// Validate if current backup file together with data to add (backupData) would surpass maximum permitted backup file size
				// .. if yes, create new backup file in next step to add remaining collection data
				backupFileSizeExceeded, err = files.WillExceedBackupFileSize(&backupData, fileNameBackupFile)
				if err != nil {
					logger.GetLogger().Error("Error in 'Backup' validating if backup file exceeding permitted files size with 'WillExceedBackupFileSize()'. File name: ", fileNameBackupFile, ". Collection name: ", collectionName, ". Error: ", err)
					return
				}
			}

			// Define backup file name depending on file size
			fileNameBackupFile, fileNumber, oldNameBackupFile, err := files.DetermineBackupFileName(fileNameBackupFile, fileNumber, collectionName, timeStampString, backupFileSizeExceeded)
			if err != nil {
				logger.GetLogger().Error("Error in 'Backup' defining file name depending on size for collection: ", collectionName, ". File number: ", fileNumber, ". File name: ", fileNameBackupFile, ". Error: ", err)
				return
			}

			// If new .csv backup file created and used to add remaining database documents due to reached size limit of backup file, add/list full/former backup file to meta data file listing
			if backupFileSizeExceeded {
				// Add entry to meta file
				timeFinalizedBackupFile := date.TimeStamp()
				err = services.AddMetaEntry(timeStamp, timeFinalizedBackupFile, folderPathBackup, fileNameMeta, oldNameBackupFile, collectionName, config.NameDatabase)
				if err != nil {
					logger.GetLogger().Error("Error in 'Backup' adding meta entry to csv for collection: ", collectionName, ". File name of meta: ", fileNameMeta, " with backup file name: ", fileNameBackupFile, " Error: ", err)
					return
				}
			}

			// Get entire name for file path and subfolder of backup file
			filePath := files.GetFilePath(fileNameBackupFile, timeStampString)

			// Create csv file for backup if not exist
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				// File does not exist, create it
				file, err := os.Create(filePath)
				if err != nil {
					panic(err)
				}

				defer file.Close()
			}

			// Write new backup data (mongo documents) to csv
			err = data.WriteBSONToCSV(backupData, filePath)
			if err != nil {
				logger.GetLogger().Error("Error in 'Backup' writing data to csv for collection: ", collectionName, " with file name: ", fileNameBackupFile, " Error: ", err)
				return
			}

			// If last iteration step for this collection add information of current and finalized backup file to meta data listing
			if i == iterations-1 {
				timeFinalizedBackupFile := date.TimeStamp()
				err = services.AddMetaEntry(timeStamp, timeFinalizedBackupFile, folderPathBackup, fileNameMeta, fileNameBackupFile, collectionName, config.NameDatabase)
				if err != nil {
					logger.GetLogger().Error("Error in 'Backup' adding meta entry to csv for collection: ", collectionName, ". File name of meta: ", fileNameMeta, " with backup file name: ", fileNameBackupFile, " Error: ", err)
					return
				}

			}

		}

	}

	/////////////////////////////////////////////////////////////////////////////////////////
	// Manage Storage Backup
	////////////////////////
	err = ManageStorages(folderPathBackup, fileNameMeta, awsClientConfig, bucketName)
	// Upload all backup files and meta data file to virtual AWS S3 folder path (folderPathBackup)
	// Manage Circular Buffer
	err = services.UploadBackupsAwsS3(folderPathBackup, fileNameMeta, awsClientConfig, bucketName)
	if err != nil {
		logger.GetLogger().Error("Error in 'Backup' applying 'UploadBackupsAwsS3()'. Error: ", err)
		return
	}

	// Manage local backups
	err = services.ManageBackupsLocally()
	if err != nil {
		logger.GetLogger().Error("Error in 'Backup' applying 'ManageBackupsLocally()'. Error: ", err)
		return
	}

	if config.SendEmailNotifications == true {
		// Setup Email client dependency
		emailMethods, err := email.GetEmailMethods(emailClientConfig)
		err = emailMethods.MethodInterface.SendEmailBackupSuccess(timeStamp, bucketName, folderPathBackup)
		if err != nil {
			logger.GetLogger().Error("Error in 'Backup' applying 'SendEmailBackupSuccess()'. Error: ", err)
			return
		}
	}
	logger.GetLogger().Infof("Backup successful. Date: %s. S3 folder path: %s. Bucket name: %s. Meta file name: %s", timeStampString, folderPathBackup, bucketName, fileNameMeta)
}
