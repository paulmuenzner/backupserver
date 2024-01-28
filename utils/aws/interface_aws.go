package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ///////////////////////////////////////////////////////////
// Setup of Dependency Injection for AWS S3 Client Methods
// /////////////////////
type S3Methods interface {
	UploadFile(bucketName string, objectKey string, filePath string) error
	DeleteObjects(bucketName string, objectKeys []string) error
	ListFolderNamesS3(bucketName, folderPrefix string) ([]string, error)
	DeleteFolderContents(bucketName, folderPrefix string) error
	BucketExists(bucketName string) (bool, error)
}

type S3Client struct {
	Client *s3.Client // Requires AWS SDK setup for actual usage
}

type AwsClientConfigData struct {
	AwsRegion      string
	AwsAccessKeyId string
	AwsSecretKey   string
}

type AwsMethodInterface struct {
	MethodInterface S3Methods
}

func NewAwsMethodInterface(s3Client *S3Client) *AwsMethodInterface {
	return &AwsMethodInterface{MethodInterface: s3Client}
}

func GetAwsMethods(awsClientConfig *AwsClientConfigData) (awsClientMethods *AwsMethodInterface, err error) {
	// Setup AWS S3 client dependency
	client, err := NewAwsClient(awsClientConfig)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create S3 client in 'AwsProductionClient()' with 'NewAwsClient()'. Error: %v", err)
	}
	awsClientMethods = NewAwsMethodInterface(client)
	return awsClientMethods, nil
}
