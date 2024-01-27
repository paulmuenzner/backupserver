package mongoDB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (client *DatabaseClient) CountDocumentsInMongo(databaseName string, collection string) (int, error) {
	var result int

	// Create a session for the database
	session, err := client.MongoDB.StartSession()
	if err != nil {
		return 0, err
	}
	defer session.EndSession(context.Background())

	// Select the database and collection
	db := client.MongoDB.Database(databaseName)
	col := db.Collection(collection)

	// Insert the data into the collection
	count, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, err
		} else {
			return 0, fmt.Errorf("Error in 'CountDocumentsInMongo()' when counting documents in collection '%s' of database '%s' Error: %v", collection, databaseName, err)
		}
	}
	result = int(count)
	return result, nil
}
