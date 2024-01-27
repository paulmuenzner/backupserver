package csv

import (
	"encoding/csv"
	"fmt"
)

// Retrieve header from csv reader
func GetCsvRow(reader *csv.Reader) (header []string, err error) {
	// Read the CSV headers to get the field names
	header, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("Error reading header in 'GetCsvRow': %v", err)
	}
	return header, nil
}
