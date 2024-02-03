package aws

import (
	"context"
	"fmt"
	"os"
	"strings"

	files "github.com/paulmuenzner/backupserver/utils/files"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

// /////////////////////////////////////////////////////////////////////
// //// UPLOADER
// Upload object to S3
func (client *S3Client) UploadFile(bucketName string, objectKey string, filePath string) error {
	// Return existing function parameter if local file path not existing
	exists, err := files.LocalFileExists(filePath)
	if err != nil {
		return fmt.Errorf("Error validating if upload file path exists in 'UploadFile' using 'LocalPathExists()'. File path %s. Error: %v", filePath, err)
	} else if !exists {
		return fmt.Errorf("Couldn't find file with file path '%s' in 'UploadFile' with 'GetFileSizeByPath()'.", filePath)
	}

	// Define size of local file to upload
	fileSize, err := files.GetFileSizeByPath(filePath)
	if err != nil {
		return fmt.Errorf("Couldn't define file size for file path '%s' in 'UploadFile' with 'GetFileSizeByPath()'. Error: %v", filePath, err)
	}

	// If file size larger than 11MB stream file in chunks with uploadLargeObjectToS3().
	// ! Minimum file size 5MB to be able to use uploadLargeObjectToS3() according to AWS
	if fileSize < 11*1024*1024 {
		err := client.uploadSmallObjectToS3(bucketName, objectKey, filePath)
		if err != nil {
			return fmt.Errorf("Error uploading file in 'UploadFile' with 'uploadSmallObjectToS3' of file path '%s' Error: %v", filePath, err)
		}
	} else {
		err := client.uploadLargeObjectToS3(bucketName, objectKey, filePath)
		if err != nil {
			return fmt.Errorf("Error uploading file in 'UploadFile' with 'uploadLargeObjectToS3' of file path '%s' Error: %v", filePath, err)
		}
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////////
// Upload files breaks large data into parts and uploads the parts concurrently
// ////////////////////////////////////////////////////////////////////////////
func (awsS3 *S3Client) uploadLargeObjectToS3(bucketName, objectKey, filePath string) error {
	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error in 'uploadLargeObjectToS3' opening file %s: %v", filePath, err)
	}
	defer file.Close()

	// Define chunk sizes
	var partMiBs int64 = 10
	uploader := manager.NewUploader(awsS3.Client, func(u *manager.Uploader) {
		u.PartSize = partMiBs * 1024 * 1024
	})

	// Stream the file in chunks directly to the uploader
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file, // Use the file directly as the input stream
	})

	if err != nil {
		return fmt.Errorf("Couldn't upload large object in 'uploadLargeObjectToS3' with 'uploader.Upload()' to bucket %v with object key:%v. Here's why: %v",
			bucketName, objectKey, err)
	}

	return err
}

// ////////////////////////////////////////////////////////
// Upload file
// ///////////
func (awsS3 *S3Client) uploadSmallObjectToS3(bucketName, objectKey, fileName string) error {
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Couldn't open file %s to upload. Error in 'uploadSmallObjectToS3' from 'os.Open(fileName)': %v", fileName, err)
	} else {
		defer file.Close()
		_, err = awsS3.Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})
		if err != nil {
			return fmt.Errorf("Couldn't upload file %s to bucket %v with object key:%v. Error in 'uploadSmallObjectToS3' from 'awsS3.Client.PutObject()': %v", fileName, bucketName, objectKey, err)
		}
	}
	return err
}

// ////////////////////////////////////////////////////////////
// Validate if S3 bucket exists
// ////////////////////////////////
func (awsS3 *S3Client) BucketExists(bucketName string) (bool, error) {
	_, err := awsS3.Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	exists := true
	if err != nil {
		err = fmt.Errorf("Either no access to bucket %s or another error determined in 'BucketExists' with 'HeadBucket()'. Error: %v", bucketName, err)
		exists = false
	}

	return exists, err
}

// ////////////////////////////////////////////////////////////
// ////////////////////////////////////////////////////////////
// Delete object from aws S3 bucket
func (awsS3 *S3Client) DeleteObjects(bucketName string, objectKeys []string) error {

	var objectIds []types.ObjectIdentifier
	for _, key := range objectKeys {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}
	_, err := awsS3.Client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{Objects: objectIds},
	})
	if err != nil {
		return fmt.Errorf("Couldn't delete objects from bucket %v. Here's why: %v", bucketName, err)
	}
	return err
}

// //////////////////////////////////////////////////////////////////////////////////////
// DeleteFolderContents recursively deletes all objects and subfolders within a folder
// ////////////////////
func (awsS3 *S3Client) DeleteFolderContents(bucketName, folderPrefix string) error {
	for {
		// List objects with pagination
		result, err := awsS3.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket:    aws.String(bucketName),
			Prefix:    aws.String(folderPrefix),
			Delimiter: aws.String("/"), // Use delimiter to group objects by common prefixes
		})
		if err != nil {
			return fmt.Errorf("Error listing objects in S3 path with 'ListObjectsV2()' in 'DeleteFolderContents'. Bucket: %s. Folder prefix: %s. Error: %v", bucketName, folderPrefix, err)
		}

		// Delete objects
		for _, object := range result.Contents {
			_, err := awsS3.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
				Bucket: aws.String(bucketName),
				Key:    object.Key,
			})
			if err != nil {
				return fmt.Errorf("Error deleting objects in S3 path with 'DeleteObject()' in 'DeleteFolderContents'. Bucket: %s. Folder prefix: %s. Object key: %s. Error: %v", bucketName, folderPrefix, *object.Key, err)
			}
		}

		// Recursively delete subfolders
		for _, commonPrefix := range result.CommonPrefixes {
			err := awsS3.DeleteFolderContents(bucketName, *commonPrefix.Prefix)
			if err != nil {
				return fmt.Errorf("Error listing objects in S3 path with 'ListObjectsV2()' in 'DeleteFolderContents'. Bucket: %s. Folder prefix: %s. Error: %v", bucketName, folderPrefix, err)
			}
		}

		// Check for more results
		if *result.IsTruncated {
			folderPrefix = *result.NextContinuationToken // Use continuation token for the next page
		} else {
			break // No more results
		}
	}

	return nil
}

// ////////////////////////////////////////////////////////////
// Check if object in S3 bucket exists
// ///////////////////////////////////
func (awsS3 *S3Client) S3ObjectExists(key, bucketName string) (bool, error) {
	// Create a HeadObjectInput with the specified key and bucket
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	// Execute the HeadObject operation
	_, err := awsS3.Client.HeadObject(context.TODO(), input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound":
				return false, nil
			default:
				return false, err
			}
		}
		return false, err
	}
	return true, nil
}

// ////////////////////////////////////////////////////////////
// List all virtual folders inside a virtual S3 folder (folderPrefix)
// ///////////////////////////////////
func (awsS3 *S3Client) ListFolderNamesS3(bucketName, folderPrefix string) ([]string, error) {
	var folderNames []string

	for {
		// List objects with pagination
		result, err := awsS3.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket:    aws.String(bucketName),
			Prefix:    aws.String(folderPrefix),
			Delimiter: aws.String("/"), // Use delimiter to group objects by common prefixes
		})
		if err != nil {
			return nil, err
		}

		// Extract folder names from CommonPrefixes
		for _, commonPrefix := range result.CommonPrefixes {
			folderName := strings.TrimSuffix(*commonPrefix.Prefix, "/") // Remove trailing slash
			folderNames = append(folderNames, folderName)
		}

		// Check for more results
		if *result.IsTruncated {
			folderPrefix = *result.NextContinuationToken // Use continuation token for the next page
		} else {
			break // No more results
		}
	}

	return folderNames, nil
}
