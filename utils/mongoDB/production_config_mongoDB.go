package mongoDB

import (
	"fmt"

	"github.com/paulmuenzner/backupserver/config"
	envHandler "github.com/paulmuenzner/backupserver/utils/env"
)

// Retrieve configuration data (eg.host, port, etc.) from .env file for production settings only
// Base parameter for dependency injection of mongodb client (production)
func MongoDBClientConfig() (mongoDBClientConfig *MongoDBClientConfigData, err error) {
	// Retrieve .env values by keys provided in config file

	// Scheme
	mongodbScheme, err := envHandler.GetEnvValue(config.MongoDatabaseSchemeEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for MongoDB scheme in 'MongoDBClientConfig()'. Env key: %s. No default value has been employed. Error: %v", config.MongoDatabaseSchemeEnv, err)
	}

	// USERNAME
	mongodbUsername, err := envHandler.GetEnvValue(config.MongoDatabaseUsernameEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for MongoDB smtp user name in 'MongoDBClientConfig()'. Env key: %s. No default value has been employed. Error: %v", config.MongoDatabaseUsernameEnv, err)
	}

	// PASSWORD
	mongodbPassword, err := envHandler.GetEnvValue(config.MongoDatabasePasswordEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for MongoDB smtp password in 'MongoDBClientConfig()'. Env key: %s. No default value has been employed. Error: %v", config.MongoDatabasePasswordEnv, err)
	}

	// HOST
	mngodbHost, err := envHandler.GetEnvValue(config.MongoDatabaseHostdEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for MongoDB smtp port in 'MongoDBClientConfig()'. Env key: %s. No default value has been employed. Error: %v", config.MongoDatabaseHostdEnv, err)
	}

	// NAME DATABASE
	mngodbDatabaseName, err := envHandler.GetEnvValue(config.MongoDatabaseNameEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for MongoDB smtp port in 'MongoDBClientConfig()'. Env key: %s. No default value has been employed. Error: %v", config.MongoDatabaseNameEnv, err)
	}

	// PORT
	mngodbPort, err := envHandler.GetEnvValue(config.MongoDatabasePortEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for MongoDB smtp port in 'MongoDBClientConfig()'. Env key: %s. No default value has been employed. Error: %v", config.MongoDatabasePortEnv, err)
	}

	// Configure MongoDBClientConfigData structure
	mongodbClientConfig := &MongoDBClientConfigData{
		Scheme:   mongodbScheme,
		Username: mongodbUsername,
		Password: mongodbPassword,
		Host:     mngodbHost,
		Port:     mngodbPort,
		Database: mngodbDatabaseName,
	}

	return mongodbClientConfig, nil
}
