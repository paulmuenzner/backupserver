package aws

import (
	"backupserver/config"
	envHandler "backupserver/utils/env"
	"fmt"
)

// Provide aws S3 connection configuration
type ProductionConfigType struct {
	AwsClientConfig *ClientConfig
	BucketName      string
	Err             error
}

func AwsS3ProductionConfig() (prodConfigOutput *ProductionConfigType) {
	// Data provided in config file
	awsRegion, err := envHandler.GetEnvValue(config.S3RegionEnvProd, "")
	if err != nil {
		return &ProductionConfigType{AwsClientConfig: nil, BucketName: "", Err: fmt.Errorf("Cannot retrieve .env value for aws region in 'awsConfig()'. Env key: %s. No default value used. Error: %v", config.S3RegionEnvProd, err)}
	}

	awsAccessKeyId, err := envHandler.GetEnvValue(config.S3AccessKeyEnvProd, "")
	if err != nil {
		return &ProductionConfigType{AwsClientConfig: nil, BucketName: "", Err: fmt.Errorf("Cannot retrieve .env value for aws access key in 'awsConfig()'. Env key: %s. No default value used. Error: %v", config.S3AccessKeyEnvProd, err)}
	}

	awsSecretKey, err := envHandler.GetEnvValue(config.S3SecretKeyEnvProd, "")
	if err != nil {
		return &ProductionConfigType{AwsClientConfig: nil, BucketName: "", Err: fmt.Errorf("Cannot retrieve .env value for aws secret key in 'awsConfig()'. Env key: %s. No default value used. Error: %v", config.S3SecretKeyEnvProd, err)}
	}

	awsClientConfig := &ClientConfig{AwsRegion: awsRegion,
		AwsAccessKeyId: awsAccessKeyId,
		AwsSecretKey:   awsSecretKey}

	// Bucket
	// bucketName := os.Getenv(config.S3BucketEnvProd)
	bucketName, err := envHandler.GetEnvValue(config.S3BucketEnvProd, "")
	if err != nil {
		return &ProductionConfigType{AwsClientConfig: nil, BucketName: "", Err: fmt.Errorf("Cannot retrieve .env value for Mongo URI in 'main.go'. Env key: %s. No default value used. Error: %v", config.S3BucketEnvProd, err)}
	}

	return &ProductionConfigType{AwsClientConfig: awsClientConfig, BucketName: bucketName, Err: nil}
}
