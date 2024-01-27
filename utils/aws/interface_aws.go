package aws

import "github.com/aws/aws-sdk-go-v2/service/s3"

/////////////////////////////////////////////////////////////
// Setup of Dependency Injection for AWS S3 Client Methods
///////////////////////
type S3Methods interface {
	UploadFile(bucketName string, objectKey string, filePath string) error
	DeleteObjects(bucketName string, objectKeys []string) error
	ListFolderNamesS3(bucketName, folderPrefix string) ([]string, error)
	DeleteFolderContents(bucketName, folderPrefix string) error
	BucketExists(bucketName string) (bool, error)
}

type AWSS3 struct {
	Client *s3.Client // Requires AWS SDK setup for actual usage
}

type ClientConfig struct {
	AwsRegion      string
	AwsAccessKeyId string
	AwsSecretKey   string
}

type MethodConfig struct {
	S3Client S3Methods
}

func NewClientBasics(s3Client *AWSS3) *MethodConfig {
	return &MethodConfig{S3Client: s3Client}
}
