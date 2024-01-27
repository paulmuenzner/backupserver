package bson

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func BsonSizeInBytes(x []bson.M) (int64, error) {
	// Marshal the BSON document to bytes

	data, err := json.Marshal(x)
	if err != nil {
		return 0, fmt.Errorf("Cannot determine bson size in 'BsonSizeInBytes'. Error: %v", err)
	}

	// Get the size in bytes
	sizeInBytes := int64(len(data))

	return sizeInBytes, nil
}
