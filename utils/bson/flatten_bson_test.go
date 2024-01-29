package bson

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFlattenBSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    bson.M
		parent   string
		sep      string
		expected map[string]interface{}
	}{
		{
			name: "Flatten Simple BSON Map",
			input: bson.M{
				"key1": "value1",
				"key2": 42,
				"key3": bson.M{"nestedKey": "nestedValue"},
			},
			parent:   "",
			sep:      "_",
			expected: map[string]interface{}{"key1": "value1", "key2": 42, "key3_nestedKey": "nestedValue"},
		},
		{
			name: "Flatten Nested BSON Map",
			input: bson.M{
				"topKey": bson.M{
					"nestedKey1": "value1",
					"nestedKey2": bson.M{"deepKey": "deepValue"},
				},
				"otherKey": "otherValue",
			},
			parent:   "",
			sep:      "-",
			expected: map[string]interface{}{"topKey-nestedKey1": "value1", "topKey-nestedKey2-deepKey": "deepValue", "otherKey": "otherValue"},
		},
		{
			name: "Flatten BSON Map with Arrays",
			input: bson.M{
				"arrayKey": []interface{}{
					bson.M{"key1": "value1"},
					42,
					"stringValue",
				},
				"otherKey": "otherValue",
			},
			parent:   "",
			sep:      "|",
			expected: map[string]interface{}{"arrayKey|0|key1": "value1", "arrayKey|1": 42, "arrayKey|2": "stringValue", "otherKey": "otherValue"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FlattenBSON(tc.input, tc.parent, tc.sep)
			assert.Equal(t, tc.expected, result)
		})
	}
}
