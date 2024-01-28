package aws

import (
	"backupserver/config"
	envHandler "backupserver/utils/env"
	"fmt"
)

// Retrieve configuration data (eg. aws region, access key) from .env file for production settings only
// Base parameter for dependency injection of aws client (production)
func AwsS3ProductionConfig() (S3ClientConfig *AwsClientConfigData, bucketName string,
	err error) {
	// Retrieve .env values by keys provided in config file

	// AWS REGION
	awsRegion, err := envHandler.GetEnvValue(config.S3RegionEnv, "")
	if err != nil {
		return nil, "", fmt.Errorf("Cannot retrieve .env value for aws region in 'AwsS3ProductionConfig()'. Env key: %s. No default value has been employed. Error: %v", config.S3RegionEnv, err)
	}

	// AWS ACCESS KEY
	awsAccessKeyId, err := envHandler.GetEnvValue(config.S3AccessKeyEnv, "")
	if err != nil {
		return nil, "", fmt.Errorf("Cannot retrieve .env value for aws access key in 'AwsS3ProductionConfig()'. Env key: %s. No default value has been employed. Error: %v", config.S3AccessKeyEnv, err)
	}

	// AWS SECRET KEY
	awsSecretKey, err := envHandler.GetEnvValue(config.S3SecretKeyEnv, "")
	if err != nil {
		return nil, "", fmt.Errorf("Cannot retrieve .env value for aws secret key in 'AwsS3ProductionConfig()'. Env key: %s. No default value has been employed. Error: %v", config.S3SecretKeyEnv, err)
	}

	// Configure AwsClientConfigData structure
	awsClientConfig := &AwsClientConfigData{AwsRegion: awsRegion,
		AwsAccessKeyId: awsAccessKeyId,
		AwsSecretKey:   awsSecretKey}

	// S3 BUCKET
	bucketName, err = envHandler.GetEnvValue(config.S3BucketEnv, "")
	if err != nil {
		return nil, "", fmt.Errorf("Cannot retrieve .env value for Mongo URI in 'AwsS3ProductionConfig()'. Env key: %s. No default value has been employed. Error: %v", config.S3BucketEnv, err)
	}

	return awsClientConfig, bucketName, nil
}
