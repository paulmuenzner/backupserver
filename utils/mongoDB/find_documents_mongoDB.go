package mongoDB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (client *DatabaseClient) FindDocumentsInMongo(databaseName string, collection string, interval int, skip int, result interface{}) error {

	// Create a session for the database
	session, err := client.MongoDB.StartSession()
	if err != nil {
		return fmt.Errorf("Error starting session with mongo client using 'StartSession()' in 'FindDocumentsInMongo'. Error: %v", err)
	}
	defer session.EndSession(context.Background())

	// Select the database and collection
	db := client.MongoDB.Database(databaseName)
	col := db.Collection(collection)

	// Retrieving data from collection
	documents, err := col.Find(context.TODO(), bson.D{}, options.Find().SetLimit(int64(interval)).SetSkip(int64(skip)))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("Error retrieving data from collection with 'Find()' in 'FindDocumentsInMongo'. Collection: %s. Interval: %d. Skip: %d. Error: %v", collection, interval, skip, err)
		} else {
			return fmt.Errorf("Error when looping collections in 'FindDocumentsInMongo' from collection '%s' of database '%s' Error: %v", collection, databaseName, err)
		}
	}
	defer documents.Close(context.Background())

	// Decode the results into the provided result interface
	if err := documents.All(context.Background(), result); err != nil {
		return fmt.Errorf("Error when decoding documents in 'FindDocumentsInMongo' from collection '%s' of database '%s' Error: %v", collection, databaseName, err)
	}

	return nil
}
