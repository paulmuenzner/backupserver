package convert

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func ConvertArrayToMap(arr []bson.M) map[string]interface{} {
	result := make(map[string]interface{})

	for i, item := range arr {
		// Assuming you want to use the index as the key
		key := fmt.Sprintf("item%d", i)

		// Convert each bson.M item to a map[string]interface{}
		itemMap := convertBSONMToMap(item)

		// Add the converted item to the result map
		result[key] = itemMap
	}

	return result
}

func convertBSONMToMap(bsonM bson.M) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range bsonM {
		switch nested := value.(type) {
		case bson.M:
			// If the value is a nested bson.M, recursively convert it
			result[key] = convertBSONMToMap(nested)
		default:
			// Otherwise, use the value as is
			result[key] = value
		}
	}

	return result
}
