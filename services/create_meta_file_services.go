package services

import (
	"fmt"
	"os"

	"github.com/paulmuenzner/backupserver/config"
	files "github.com/paulmuenzner/backupserver/utils/files"
)

// Create csv file for meta data of all created backup files
func CreateMetaDataFile(folderPath, fileNameMeta string) (err error) {
	// Create meta file
	filePath := folderPath + "/" + fileNameMeta
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error creating file with 'os.Create()' in 'CreateMetaDataFile'. File name: %s. Error: %v", filePath, err)
	}

	defer file.Close()

	// Write headers
	var headers []string = config.MetaFileHeaders
	err = files.WriteHeadersToCsvFile(file, headers)
	if err != nil {
		return fmt.Errorf("Error writing header to csv file with 'WriteHeadersToCsvFile()' in 'CreateMetaDataFile'. File name: %s. Error: %v", filePath, err)
	}

	return nil
}
