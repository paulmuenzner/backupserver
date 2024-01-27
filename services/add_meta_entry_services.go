package services

import (
	"backupserver/config"
	files "backupserver/utils/files"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Add row of meta data for one csv file
func AddMetaEntry(dateStartBackup time.Time, dateFinalizedBackupFile time.Time, folderPathBackup, fileNameMeta, fileNameBackupFile, collectionName, databaseName string) (err error) {
	filePathMeta := filepath.Join(folderPathBackup, fileNameMeta)
	filePathBackupFile := filepath.Join(folderPathBackup, fileNameBackupFile+".csv")

	/////////////////////////////////////////////////////////////////
	// Open Meta File
	metaFile, err := os.OpenFile(filePathMeta, os.O_RDWR, 0644)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("Permission denied to open file: %s in 'AddMetaEntry'. Error: %v", filePathMeta, err)
		} else {
			return fmt.Errorf("Error opening file path: %s in 'AddMetaEntry'. Error: %v", filePathMeta, err)
		}
	}
	defer metaFile.Close()

	// Check if meta file header is valid
	err = files.ValidateCSVHeaders(metaFile, config.MetaFileHeaders)
	if err != nil {
		return fmt.Errorf("Detected in 'AddMetaEntry' that meta file header is not valid for meta file %s. Error: %v", filePathMeta, err)
	}

	// Initialize CSV writer
	metaWriter := csv.NewWriter(metaFile)

	// Open Backup File
	backupFile, err := os.OpenFile(filePathBackupFile, os.O_RDWR, 0644)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("Permission denied to open backup file: %s in 'AddMetaEntry'. Error: %v", fileNameBackupFile, err)
		} else {
			return fmt.Errorf("Error opening backup file path: %s in 'AddMetaEntry'. Error: %v", fileNameBackupFile, err)
		}
	}
	defer backupFile.Close()

	// Get file size of backup file
	sizeBackupFile, err := files.GetFileSizeByDescriptor(backupFile)
	if err != nil {
		return fmt.Errorf("Error determining backup file size with 'GetFileSizeByDescriptor()' in 'AddMetaEntry'. File name: %s. Error: %v\n", filePathBackupFile, err)
	}

	// Create a new row of data
	dataRow := config.RowTypesMeta{
		CollectionName:    collectionName,
		FolderPath:        folderPathBackup,
		FileName:          fileNameBackupFile,
		SizeInBytes:       sizeBackupFile,
		DatabaseName:      databaseName,
		DateStartBackup:   dateStartBackup,
		DateFinalizedFile: dateFinalizedBackupFile,
	}

	// Write the data row to the CSV file
	if err := files.AddOneMetaRow(metaWriter, dataRow); err != nil {
		return fmt.Errorf("Error adding one row to csv with 'AddOneMetaRow()' in 'AddMetaEntry'. File name: %s. Error: %v\n", filePathBackupFile, err)
	}

	// Flush and close CSV writer
	metaWriter.Flush()
	if err := metaWriter.Error(); err != nil {
		return fmt.Errorf("Error flushing and closing csv writer with 'AddOneMetaRow()' in 'AddMetaEntry'. File name: %s. Error: %v\n", filePathBackupFile, err)
	}

	return nil
}
