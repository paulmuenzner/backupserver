package bson

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBsonSizeInBytes(t *testing.T) {
	testCases := []struct {
		name     string
		input    []bson.M
		expected int64
		err      error
	}{
		{
			name:     "Empty Array",
			input:    []bson.M{},
			expected: 2, // Expected size for an empty array (considering BSON format),
			err:      nil,
		},
		{
			name: "Array with One Document",
			input: []bson.M{
				{"key1": "value1", "key2": 42},
			},
			expected: 29, // Expected size for a simple BSON document (considering BSON format),
			err:      nil,
		},
		{
			name: "Array with Nested Documents",
			input: []bson.M{
				{"key1": "value1", "key2": bson.M{"nestedKey": "nestedValue"}},
				{"key3": 42},
			},
			expected: 66, // Expected size for documents with nesting and array elements,
			err:      nil,
		},
		{
			name: "Array with Complex Nested Documents",
			input: []bson.M{
				{"key1": "value1", "key2": bson.M{"nestedKey1": "nestedValue1", "nestedKey2": bson.M{"deepKey": "deepValue"}}},
				{"key3": 42},
				{"key4": bson.M{"nestedKey3": "nestedValue3"}},
			},
			expected: 144, // Expected size for complex nested documents and array elements,
			err:      nil,
		},
		// Add more test cases as needed
	}

	var wg sync.WaitGroup

	for _, tc := range testCases {
		wg.Add(1)
		go func(tc struct {
			name     string
			input    []bson.M
			expected int64
			err      error
		}) {
			defer wg.Done()

			result, err := BsonSizeInBytes(tc.input)

			assert.Equal(t, tc.expected, result)

			if tc.err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.err.Error())
			}
		}(tc)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
