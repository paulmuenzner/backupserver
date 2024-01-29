package bson

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func FlattenBSON(entry bson.M, parentKey string, sep string) map[string]interface{} {
	flat := make(map[string]interface{})
	for k, v := range entry {
		var key string
		if parentKey == "" {
			key = fmt.Sprintf("%s%s", parentKey, k)
		} else {
			key = fmt.Sprintf("%s%s%s", parentKey, sep, k)
		}

		switch value := v.(type) {
		case bson.M:
			flatMap := FlattenBSON(value, key, sep)
			for fk, fv := range flatMap {
				flat[fk] = fv
			}
		case []interface{}:
			for i, item := range value {
				arrayKey := fmt.Sprintf("%s%s%d", key, sep, i)
				if reflect.TypeOf(item).Kind() == reflect.Map {
					flatMap := FlattenBSON(item.(bson.M), arrayKey, sep)
					for fk, fv := range flatMap {
						flat[fk] = fv
					}
				} else {
					flat[arrayKey] = item
				}
			}
		default:
			flat[key] = value
		}
	}
	return flat
}
