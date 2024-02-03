package backup

import (
	"fmt"
	"math"
	"os"
	"time"

	services "github.com/paulmuenzner/backupserver/services"
	data "github.com/paulmuenzner/backupserver/utils/csv"
	date "github.com/paulmuenzner/backupserver/utils/date"
	files "github.com/paulmuenzner/backupserver/utils/files"
	mongoDB "github.com/paulmuenzner/backupserver/utils/mongoDB"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateBackupFiles(databaseClientSetup *mongoDB.MongoDBMethodInterface, databaseName string, timeStamp time.Time, folderPathBackup, fileNameMeta string) error {

	// Get current time stamp
	timeStampString := date.TimeStampSlug(timeStamp)

	////////////////////////////////////////////////////////////////////////////////////////
	// Fetch Names of All Collections in the Database. Result is a list of type []string.
	collections, err := databaseClientSetup.MethodInterface.GetAllCollections(databaseName)
	if err != nil {
		return fmt.Errorf("Error in 'CreateBackupFiles()' retrieving collection list with 'GetAllCollections()' for database name '%s'. Error: %v", databaseName, err)

	}

	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Loop through all database collections and create backup files in local backup folder 'folderPathBackup'
	for _, collectionName := range collections {
		// Count number of documents in MongoDB collection
		totalDocuments, err := databaseClientSetup.MethodInterface.CountDocumentsInMongo(databaseName, collectionName)
		if err != nil {
			return fmt.Errorf("Error in 'CreateBackupFiles()' utilizing 'CountDocumentsInMongo()' counting documents in collection: %s. Error: %v", collectionName, err)
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
			var backupData []bson.M
			err := databaseClientSetup.MethodInterface.FindDocumentsInMongo(databaseName, collectionName, interval, skip, &backupData)
			if err != nil {
				return fmt.Errorf("Error in 'CreateBackupFiles()' utilizing 'FindDocumentsInMongo()' retrieving documents from collection: '%s'. Skip: '%d'. Error: %v", collectionName, skip, err)
			}

			// Check if currently determined backup file is existing
			backupFileCreated, err := files.LocalFileExists(fileNameBackupFile)
			if err != nil {
				return fmt.Errorf("Error validating if file path exists in 'CreateBackupFiles()' utilizing 'LocalFileExists()'. File path '%s'. Error: %v", fileNameBackupFile, err)
			}

			var backupFileSizeExceeded bool = false
			if backupFileCreated {
				// Validate if current backup file together with data to add (backupData) would surpass maximum permitted backup file size
				// .. if yes, create new backup file in next step to add remaining collection data
				backupFileSizeExceeded, err = files.WillExceedBackupFileSize(&backupData, fileNameBackupFile)
				if err != nil {
					return fmt.Errorf("Error in 'CreateBackupFiles()' validating if backup file exceeding permitted files size with 'WillExceedBackupFileSize()'. File name: '%s'. Collection name: '%s'. Error: %v", fileNameBackupFile, collectionName, err)
				}
			}

			// Define backup file name depending on file size
			fileNameBackupFile, fileNumber, oldNameBackupFile, err := files.DetermineBackupFileName(fileNameBackupFile, fileNumber, collectionName, timeStampString, backupFileSizeExceeded)
			if err != nil {
				return fmt.Errorf("Error in 'CreateBackupFiles()' utilizing 'DetermineBackupFileName()' defining file name depending on size for collection: '%s'. File number: '%d'. Backup file name: '%s'. Error: %v", collectionName, fileNumber, fileNameBackupFile, err)
			}

			// If new .csv backup file created and used to add remaining database documents due to reached size limit of backup file, add/list full/former backup file to meta data file listing
			if backupFileSizeExceeded {
				// Add entry to meta file
				timeFinalizedBackupFile := date.TimeStamp()
				err = services.AddMetaEntry(timeStamp, timeFinalizedBackupFile, folderPathBackup, fileNameMeta, oldNameBackupFile, collectionName, databaseName)
				if err != nil {
					return fmt.Errorf("Error in 'CreateBackupFiles()' adding meta entry to csv for collection: '%s'. Meta file name: '%s'. Backup file name:  '%s'. Error: %v", collectionName, fileNameMeta, fileNameBackupFile, err)
				}
			}

			// Get entire name for file path and subfolder of backup file
			filePath := files.GetFilePath(fileNameBackupFile, timeStampString)

			// Create csv file for backup if not exist
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				// File does not exist, create it
				file, err := os.Create(filePath)
				if err != nil {
					return fmt.Errorf("Error in 'CreateBackupFiles()'. Unable to create csv backup file for collection: '%s'. Meta file name: '%s'. Backup file path: '%s'. Error: %v", collectionName, fileNameMeta, filePath, err)
				}

				defer file.Close()
			}

			// Write new backup data (mongo documents) to related csv file
			err = data.WriteBSONToCSV(backupData, filePath)
			if err != nil {
				return fmt.Errorf("Error in 'CreateBackupFiles()' writing data to csv for collection: '%s'. Backup file name: '%s'. Error: %v", collectionName, fileNameBackupFile, err)
			}

			// If last iteration step for this collection add information of current and finalized backup file to meta data listing
			if i == iterations-1 {
				timeFinalizedBackupFile := date.TimeStamp()
				err = services.AddMetaEntry(timeStamp, timeFinalizedBackupFile, folderPathBackup, fileNameMeta, fileNameBackupFile, collectionName, databaseName)
				if err != nil {
					return fmt.Errorf("Error in 'CreateBackupFiles()' adding meta entry to csv for collection '%s'. File name of meta: '%s' with backup file name: '%s'. Error: %v", collectionName, fileNameMeta, fileNameBackupFile, err)
				}

			}

		}
	}

	return nil
}
