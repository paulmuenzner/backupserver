package main

import (
	"context"
	"time"

	"github.com/paulmuenzner/backupserver/config"
	cronJobs "github.com/paulmuenzner/backupserver/cron/backup"
	aws "github.com/paulmuenzner/backupserver/utils/aws"
	email "github.com/paulmuenzner/backupserver/utils/email"
	envHandler "github.com/paulmuenzner/backupserver/utils/env"
	logger "github.com/paulmuenzner/backupserver/utils/logs"
	mongoDB "github.com/paulmuenzner/backupserver/utils/mongoDB"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig struct {
	MongoClient *mongo.Client
}

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logger.GetLogger().Errorf("Error in 'main()' loading .env file: %v", err)
		return
	}

	///////////////////////////////////////////////
	// LOGGER MIDDLEWARE ////////////
	///////////////////////////////////////////////

	logFileName := logger.GetLogFileName()
	logger.Init(logFileName)

	///////////////////////////////////////////////
	// LOGGER MIDDLEWARE ////////
	///////////////////////////////////////////////

	///////////////////////////////////////////////
	// CONNECT DATABASE MONGODB ///////////////////
	///////////////////////////////////////////////

	// Get Uniform Resource Identifier
	mongodbURI, err := mongoDB.MongoDBClientConfig()
	if err != nil {
		logger.GetLogger().Warnf("Cannot retrieve .env value for Mongo URI in 'main.go'. Default value used. Error: %v", err)
	}

	// Get database name
	databaseName, err := envHandler.GetEnvValue(config.MongoDatabaseNameEnv, "")
	if err != nil {
		logger.GetLogger().Errorf("Cannot retrieve .env value for database name (config.MongoDatabaseNameEnv) in 'main()' utilizing 'GetEnvValue()'. Env key: %s. No default value has been employed. Error: %v", config.MongoDatabaseNameEnv, err)
		return
	}

	// Connect to database
	client, err := mongoDB.ConnectToMongoDB(mongodbURI)
	if err != nil {
		logger.GetLogger().Errorf("Connection with MongoDB failed due to following error: %v", err)
		return
	}

	// Disconnect from MongoDB
	defer func() {
		if err := client.MongoDB.Disconnect(context.TODO()); err != nil {
			logger.GetLogger().Errorf("Disconnection MongoDB failed due to following error: %v", err)
			return
		}
	}()

	///////////////////////////////////////////////
	// END CONNECT DATABASE MONGODB ///////////////
	///////////////////////////////////////////////

	///////////////////////////////////////////////
	// SETUP CRON JOBS ////////////////////////////
	///////////////////////////////////////////////
	// More -> https://pkg.go.dev/github.com/robfig/cron#section-readme
	// Create a new cron scheduler
	cron := cron.New()

	// Cron Backup
	// AWS S3 client config production
	s3ClientConfig, bucketName, err := aws.AwsS3ProductionConfig()
	if err != nil {
		logger.GetLogger().Error("Error in 'Backup' retrieving aws configuration using 'awsConfig()'. Error: ", err)
		return
	}

	// Email client config production
	emailClientConfig, err := email.EmailProductionConfig()
	if err != nil {
		logger.GetLogger().Error("Error in 'Backup' retrieving aws configuration using 'awsConfig()'. Error: ", err)
		return
	}

	// Start cron job
	_, errCron1 := cron.AddFunc(config.IntervalBackup, func() {
		cronJobs.Backup(client, s3ClientConfig, emailClientConfig, bucketName, databaseName)
	}) // Use the imported function
	if errCron1 != nil {
		logger.GetLogger().Error("Error adding cron job 'Backup': ", errCron1)
		return
	}

	// Start the cron scheduler in a separate goroutine
	go func() {
		cron.Start()
		defer cron.Stop() // Stop the cron scheduler when the goroutine exits
		select {}
	}()

	///////////////////////////////////////////////
	// END SETUP CRON JOBS ////////////////////////
	///////////////////////////////////////////////

	for {
		time.Sleep(1 * time.Minute)
	}

}
