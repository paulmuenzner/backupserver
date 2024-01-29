package mongoDB

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// ///////////////////////////////////////////////////////////
// Setup of Dependency Injection for MongoDB Client Methods
// /////////////////////
type MongoDBMethods interface {
	CountDocumentsInMongo(databaseName string, collection string) (int, error)
	GetAllCollections(nameDatabase string) ([]string, error)
	FindDocumentsInMongo(databaseName string, collection string, interval int, skip int, result interface{}) error
}

type MongoDBClient struct {
	MongoDB *mongo.Client
}

type MongoDBClientConfigData struct {
	Scheme   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type MongoDBMethodInterface struct {
	MethodInterface MongoDBMethods
}

func NewMongoDBMethodInterface(mongoClient *MongoDBClient) *MongoDBMethodInterface {
	return &MongoDBMethodInterface{MethodInterface: mongoClient}
}
