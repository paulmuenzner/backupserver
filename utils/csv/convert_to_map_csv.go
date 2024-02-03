package csv

import (
	"fmt"
	"os"
	"reflect"
)

func ConvertCsvToMap(filePath string, recordType interface{}) (csvAsMap []map[string]string, err error) {
	// Open file
	metaFile, err := os.OpenFile(filePath, os.O_RDWR, 0600)
	if err != nil {
		if os.IsPermission(err) {
			return nil, fmt.Errorf("Permission denied to open file: %s in 'ConvertCsvToMap'. Error: %v", filePath, err)
		} else {
			return nil, fmt.Errorf("Error opening file path: %s in 'ConvertCsvToMap'. Error: %v", filePath, err)
		}
	}
	defer metaFile.Close()

	// Create a CSV reader
	reader := NewCsvReader(metaFile)

	// Retrieve csv header
	csvHeader, err := reader.Read() // GetCsvRow(reader)
	if err != nil {
		return nil, fmt.Errorf("Error with 'GetCsvRow()' in 'ConvertCsvToMap'. File path: %s. Error: %v", filePath, err)
	}

	// Get the type of the struct passed as a parameter
	recordTypeRef := reflect.TypeOf(recordType)

	// Validate if the recordType is a struct
	if recordTypeRef.Kind() != reflect.Struct {
		return nil, fmt.Errorf("RecordType must be a struct. Not struct in 'ConvertCsvToMap'.")
	}

	// Create a map to store the data
	csvAsMap = make([]map[string]string, 0)

	// Read the remaining rows
	for {
		row, err := GetCsvRow(reader)
		if err != nil {
			break
		}

		// Create a map for each row
		rowData := make(map[string]string)

		// Populate the map with data
		for i, value := range row {
			rowData[csvHeader[i]] = value
		}

		// Add the row map to the data slice
		csvAsMap = append(csvAsMap, rowData)
	}

	return csvAsMap, nil
}
