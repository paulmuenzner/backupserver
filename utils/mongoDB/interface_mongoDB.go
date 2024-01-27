package mongoDB

import "go.mongodb.org/mongo-driver/mongo"

/////////////////////////////////////////////////////////////
// Setup of Dependency Injection for MongoDB Client Methods
///////////////////////
type MongoDBMethods interface {
	CountDocumentsInMongo(databaseName string, collection string) (int, error)
	GetAllCollections(nameDatabase string) ([]string, error)
	FindDocumentsInMongo(databaseName string, collection string, interval int, skip int, result interface{}) error
}

type DatabaseClient struct {
	MongoDB *mongo.Client
}

type MethodConfig struct {
	MethodInterface MongoDBMethods
}

func MongoClientBasics(mongoClient *DatabaseClient) *MethodConfig {
	return &MethodConfig{MethodInterface: mongoClient}
}
