package services

import (
	config "backupserver/config"
	aws "backupserver/utils/aws"
	csvHandler "backupserver/utils/csv"
	"fmt"
)

func UploadBackupsAwsS3(folderPathBackup, metaFileName string, awsClientConfig *aws.ClientConfig, bucketName string) error {

	// Create an instance of RowTypesMeta
	rowMetaInstance := config.RowTypesMeta{}

	// Open related meta file retrieving information of all backup files
	filePathMeta := folderPathBackup + "/" + metaFileName
	metaDataAsMap, err := csvHandler.ConvertCsvToMap(filePathMeta, rowMetaInstance)
	if err != nil {
		return fmt.Errorf("Couldn't convert csv to map for file path %s in 'UploadBackupsAwsS3' with 'ConvertCsvToMap()'. Error: %v\n", filePathMeta, err)
	}

	// Setup AWS S3 client dependency
	client, err := aws.NewS3Client(awsClientConfig)
	if err != nil {
		return fmt.Errorf("Couldn't create S3 client in 'UploadBackupsAwsS3' with 'NewS3Client(awsClientConfig)'. Error: %v\n", err)
	}
	basics := aws.NewClientBasics(client)

	// Validate if bucket accessible
	bucketExist, err := basics.S3Client.BucketExists(bucketName)
	if !bucketExist {
		return fmt.Errorf("S3 Bucket %s not found in 'UploadBackupsAwsS3' with 'BucketExists(bucketName)'. Error:  %v\n", bucketName, err)
	}
	if err != nil {
		return fmt.Errorf("Error in 'UploadBackupsAwsS3' validating if bucket '%s' accessible with 'BucketExists(bucketName)'. Error:  %v\n", bucketName, err)
	}

	// Loop through each backup file row listed in meta file
	for _, value := range metaDataAsMap {
		filePath := value["folder_path"] + "/" + value["file_name"] + ".csv"

		err = basics.S3Client.UploadFile(bucketName, filePath, filePath)

		if err != nil {
			return fmt.Errorf("Couldn't upload file in 'UploadBackupsAwsS3' with 'UploadFile()'. Error: %v\n", err)
		}
	}

	// Add meta data file itself to backup path on S3
	err = basics.S3Client.UploadFile(bucketName, filePathMeta, filePathMeta)
	if err != nil {
		return fmt.Errorf("Couldn't upload meta file of path '%s' in 'UploadBackupsAwsS3' with 'UploadFile()'. Error: %v\n", filePathMeta, err)
	}

	// Circular buffer S3 - Only store latest number of n backups and delete older ones if circular buffer activated
	isCircularBufferActivatedS3 := config.IsCircularBufferActivatedS3
	if isCircularBufferActivatedS3 == true {
		err := DeleteOldBackupsS3(bucketName, awsClientConfig)
		if err != nil {
			return fmt.Errorf("Error in 'Backup' applying 'DeleteOldBackups()'. Error: %v\n", err)

		}
	}

	return nil
}
