package mongoDB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func ConnectToMongoDB(mongodbURI string) (*DatabaseClient, error) {

	commandStarted := []string{}
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			commandStarted = append(commandStarted, evt.CommandName)
		},
	}

	clientOptions := options.Client().ApplyURI(mongodbURI).
		// Add your security settings here (e.g., clientOptions.SetAuth)
		SetMaxPoolSize(50).
		SetReadConcern(readconcern.Local()).
		SetWriteConcern(writeconcern.Majority()).
		SetRetryWrites(true).
		SetCompressors([]string{"zstd", "snappy"}).
		SetMonitor(cmdMonitor)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Client setup MongoDB failed in 'ConnectToMongoDB()' applying 'mongo.Connect()'. Error: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error in 'ConnectToMongoDB()' applying 'client.Ping()'. Cannot connect to MongoDB. Error: %v", err)
	}

	return &DatabaseClient{MongoDB: client}, nil
}
