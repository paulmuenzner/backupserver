package files

import (
	"fmt"

	"github.com/paulmuenzner/backupserver/config"
	bsonHandler "github.com/paulmuenzner/backupserver/utils/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WillExceedBackupFileSize(dataToAdd *[]primitive.M, filePath string) (fileTooLarge bool, err error) {
	// Determine size of currently used backup file
	existingFileSize, err := GetFileSizeByPath(filePath)
	if err != nil {
		return true, fmt.Errorf("Error in 'WillExceedBackupFileSize' determining size of file '%s' with 'GetFileSizeByPath()' documents. Error: %v", filePath, err)
	}

	// Determine size of new data (mongo documents) to add to currently used backup file
	sizeBson, err := bsonHandler.BsonSizeInBytes(*dataToAdd)
	securityBuffer := 1.1 // 10% security buffer
	if err != nil {
		return true, fmt.Errorf("Error in 'WillExceedBackupFileSize' determining size of retrieved database documents with 'BsonSizeInBytes()'. Error: %v", err)
	}

	// Determine if backup file would exceed file size if new data would be added
	totalSizeBytes := float64(existingFileSize) + float64(sizeBson)*securityBuffer
	maxFileSize := config.MaxFileSizeInBytes
	wouldExceedPermittedFileSize := totalSizeBytes > float64(maxFileSize)

	if wouldExceedPermittedFileSize {
		return true, nil
	} else {
		return false, nil
	}
}
