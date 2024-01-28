package services

import (
	config "backupserver/config"
	aws "backupserver/utils/aws"
	"backupserver/utils/helper"
	"fmt"
	"regexp"
)

// This backup server can optionally be used as a ring memory. Due to this, it is needed to keep the newest backups of number n
func DeleteOldBackupsS3(bucketName string, awsClientConfig *aws.AwsClientConfigData) error {
	// Setup AWS S3 client dependency
	awsMethods, err := aws.GetAwsMethods(awsClientConfig)
	if err != nil {
		return fmt.Errorf("Error in 'UploadBackupsAwsS3()' with 'AwsMethodInterface()'. Error:  %v", bucketName, err)
	}

	///////////////////////////////////////////////////////////////////
	// Retrieve virtual folder names inside backup folder of S3 bucket
	folderPrefix := config.FolderNameBackup + "/"
	folderNames, err := awsMethods.MethodInterface.ListFolderNamesS3(bucketName, folderPrefix)
	if err != nil {
		return fmt.Errorf("Couldn't create S3 client in 'DeleteOldBackupsS3'. Error: %v", err)

	}

	/////////////////////////////////////////////////////////////////////////////
	// Continue if more backups exist than permitted (>> config.MaxBackups)
	maxNumberBackups := config.MaxBackupsS3
	moreBackupsThanPermitted := moreBackupsThanPermitted(maxNumberBackups, folderNames)
	if moreBackupsThanPermitted {
		//////////////////////////////////////////////////////////
		// Extract date format from slice of folder names
		backupDates := extractDateFromPath(folderNames)

		//////////////////////////////////////////////////////////
		// Get old backup folder names to delete
		outdatedBackups, _ := helper.ClassifyBackupsByAge(backupDates, maxNumberBackups)

		//////////////////////////////////////////////////////////
		// Delete outdated backups
		for _, outdatedBackupFolder := range outdatedBackups {
			path := config.FolderNameBackup + "/" + outdatedBackupFolder
			err := awsMethods.MethodInterface.DeleteFolderContents(bucketName, path)
			if err != nil {
				return fmt.Errorf("Error in 'DeleteOldBackupsS3': %v", err)
			}
		}
	}

	return nil
}

// Validate if more backups in S3 bucket than permitted
func moreBackupsThanPermitted(permitted int, backups []string) bool {
	return len(backups) > permitted
}

// Format and extract date from folder paths
func extractDateFromPath(backupPaths []string) []string {
	stringSlice := make([]string, len(backupPaths))

	// Regex to match only numbers, hyphens, and underscores
	reg := regexp.MustCompile("[^\\d-_]+")

	for i, str := range backupPaths {
		modifiedStr := reg.ReplaceAllString(str, "")
		stringSlice[i] = modifiedStr
	}

	return stringSlice
}
