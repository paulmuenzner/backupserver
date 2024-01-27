package main

import (
	"backupserver/config"
	aws "backupserver/utils/aws"
	cronJobs "backupserver/utils/cron"
	envHandler "backupserver/utils/env"
	logger "backupserver/utils/logs"
	mongoDB "backupserver/utils/mongoDB"
	"context"
	"time"

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
	mongodbURI, err := envHandler.GetEnvValue("MONGO_URI", "mongodb://localhost:27017")
	if err != nil {
		logger.GetLogger().Warnf("Cannot retrieve .env value for Mongo URI in 'main.go'. Default value used. Error: %v", err)
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
	// AWS S3 client config
	prodConfigOutput := aws.AwsS3ProductionConfig()
	if prodConfigOutput.Err != nil {
		logger.GetLogger().Error("Error in 'Backup' retrieving aws configuration using 'awsConfig()'. Error: ", err)
		return
	}
	_, errCron1 := cron.AddFunc(config.IntervalBackup, func() { cronJobs.Backup(client, prodConfigOutput.AwsClientConfig, prodConfigOutput.BucketName) }) // Use the imported function
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
