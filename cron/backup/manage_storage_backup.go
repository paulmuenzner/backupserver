package backup

import (
	"fmt"

	services "github.com/paulmuenzner/backupserver/services"
	"github.com/paulmuenzner/backupserver/utils/aws"
)

func ManageStorages(folderPathBackup string, metaFileName string, awsClientConfig *aws.AwsClientConfigData, bucketName string) error {
	/////////////////////////////////////////////////////////////////////////////////////////
	// Manage Backup Storage
	////////////////////////
	//
	// AWS S3
	//
	// Upload all backup files and meta data file to virtual AWS S3 folder path (folderPathBackup)
	// Manage Circular Buffer
	err := services.UploadBackupsAwsS3(folderPathBackup, metaFileName, awsClientConfig, bucketName)
	if err != nil {
		return fmt.Errorf("Error in 'Backup' utilizing 'UploadBackupsAwsS3()'. Error: %v", err)

	}

	//
	// Local
	//
	// Manage local backups depending on configuration in '/config/base_config.go' => Delete all, keep all or circular buffer
	err = services.ManageBackupsLocally()
	if err != nil {
		return fmt.Errorf("Error in 'Backup' utilizing 'ManageBackupsLocally()'. Error: %v", err)
	}

	return nil
}
