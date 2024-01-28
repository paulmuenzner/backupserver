package files

import (
	"fmt"
	"strconv"
)

// Determine file name depending on backup file size.
// Create new backup file name if cureent file size exceeds maximum size by using base backup file name and applying subsequent numbering
func DetermineBackupFileName(fileName string, fileNumber int, collectionName string, timeStamp string, fileSizeExceeding bool) (newFileName string, newFileNumber int, oldNameBackupFile string, err error) {
	filePath := GetFilePath(fileName, timeStamp)

	// Return existing function parameter if file path not existing
	exists, err := LocalFileExists(filePath)
	if err != nil {
		return "", fileNumber, "", fmt.Errorf("Error validating if file path exists in 'DetermineBackupFileName' using 'LocalPathExists()'. File path %s. Error: %v", filePath, err)
	} else if exists == false {
		return fileName, fileNumber, fileName, nil
	}

	newFileName = fileName
	newFileNumber = fileNumber

	// Create new file name if backup file would exceed size
	if fileSizeExceeding {
		// Create new backup file name as maximum file size has been reached (collection will be written to many csv files)
		newFileName = collectionName + "_" + strconv.FormatInt(int64(fileNumber+1), 10)
		newFileNumber++
	}

	return newFileName, newFileNumber, fileName, nil
}
