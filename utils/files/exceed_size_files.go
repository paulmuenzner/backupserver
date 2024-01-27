package files

import (
	"backupserver/config"
	bsonHandler "backupserver/utils/bson"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WillExceedBackupFileSize(dataToAdd *[]primitive.M, filePath string) (fileTooLarge bool, err error) {
	// Determine size of currently used backup file
	existingFileSize, err := GetFileSizeByPath(filePath)
	if err != nil {
		return true, fmt.Errorf("Error in 'WillExceedBackupFileSize' determining size of file '%s' with 'GetFileSizeByPath()' documents. Error: %v\n", filePath, err)
	}

	// Determine size of new data (mongo documents) to add to currently used backup file
	sizeBson, err := bsonHandler.BsonSizeInBytes(*dataToAdd)
	if err != nil {
		return true, fmt.Errorf("Error in 'WillExceedBackupFileSize' determining size of retrieved database documents with 'BsonSizeInBytes()'. Error: %v\n", err)
	}

	// Determine if backup file would exceed file size if new data would be added
	totalSizeBytes := existingFileSize + sizeBson
	maxFileSize := config.MaxFileSizeInBytes
	wouldExceedPermittedFileSize := totalSizeBytes > maxFileSize

	if wouldExceedPermittedFileSize {
		return true, nil
	} else {
		return false, nil
	}
}
