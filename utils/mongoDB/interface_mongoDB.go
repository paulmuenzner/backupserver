package mongoDB

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// ///////////////////////////////////////////////////////////////////////////
// Setup interface for database repository utilizing Dependency Injection
// ///////////////////
type MongoDBRepository interface {
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
	MethodInterface MongoDBRepository
}

func NewMongoDBMethodInterface(mongoClient *MongoDBClient) *MongoDBMethodInterface {
	return &MongoDBMethodInterface{MethodInterface: mongoClient}
}
