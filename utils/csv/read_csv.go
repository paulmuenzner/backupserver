package csv

import (
	"encoding/csv"
	"os"
)

// Create csv reader
func NewCsvReader(file *os.File) (csvReader *csv.Reader) {
	// Create a CSV reader
	csvReader = csv.NewReader(file)
	return csvReader
}
