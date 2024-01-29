package mongoDB

import (
	"context"
	"fmt"

	logger "github.com/paulmuenzner/backupserver/utils/logs"

	"go.mongodb.org/mongo-driver/bson"
)

func (client *MongoDBClient) DatabaseExists(nameDatabase string) (bool, error) {
	// List all databases
	databases, err := client.MongoDB.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		return false, err
	}

	// Check if the target database is in the list
	for _, db := range databases {
		if db == nameDatabase {
			return true, nil
		}
	}

	return false, nil
}

func (client *MongoDBClient) GetAllCollections(nameDatabase string) ([]string, error) {

	// Get a handle for the database
	database := client.MongoDB.Database(nameDatabase)

	exists, err := client.DatabaseExists(nameDatabase)
	if err != nil {
		return nil, fmt.Errorf("Database of name '%v' not found in 'GetAllCollections' with 'DatabaseExists()'. Error: %v", nameDatabase, err)
	}

	if !exists {
		return nil, fmt.Errorf("Database not found in 'GetAllCollections'. Database name: %s", nameDatabase)
	}

	// List all collections in the database
	var filter bson.M = bson.M{}
	collections, err := database.ListCollectionNames(context.Background(), filter)
	if err != nil {
		logger.GetLogger().Error("Not able to retrieve all collection names in 'GetAllCollections'. Error: ", err)
		return nil, err
	}

	return collections, nil
}
