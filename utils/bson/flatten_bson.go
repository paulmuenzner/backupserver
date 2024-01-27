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

		if reflect.TypeOf(v).Kind() == reflect.Map {
			flatMap := FlattenBSON(v.(bson.M), key, sep)
			for fk, fv := range flatMap {
				flat[fk] = fv
			}
		} else {
			flat[key] = v
		}
	}
	return flat
}
