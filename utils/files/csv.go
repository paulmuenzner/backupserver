package files

import (
	"backupserver/config"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// /////////////////////////////////////////////////////////////////////////
// Add one row to CSV file
func AddOneCsvRow(writer *csv.Writer, row []string) error {
	if err := writer.Write(row); err != nil {
		return fmt.Errorf("Error adding csv row to file with 'writer.Write()' in 'AddOneCsvRow'. Error: %v", err)
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////
// Add a single row to meta CSV file using specific format
func AddOneMetaRow(writer *csv.Writer, dataRow config.RowTypesMeta) error {
	// Convert dataRow to []string
	stringSlice := []string{
		dataRow.CollectionName,
		dataRow.FolderPath,
		dataRow.FileName,
		strconv.FormatInt(dataRow.SizeInBytes, 10),
		dataRow.DatabaseName,
		dataRow.DateStartBackup.Format("2006-01-02 15:04:05.000000"),
		dataRow.DateFinalizedFile.Format("2006-01-02 15:04:05.000000"),
	}

	// Write row to csv file
	if err := writer.Write(stringSlice); err != nil {
		return fmt.Errorf("Error adding meta row with 'writer.Write()' in 'AddOneMetaRow'. Error: %v", err)
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////
// Write headers (using slice of strings) to csv file
func WriteHeadersToCsvFile(file *os.File, headers []string) error {
	writer := csv.NewWriter(file)

	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("Error writing header header to file with 'writer.Write()' in 'WriteHeadersToCsvFile'. Error: %v", err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return fmt.Errorf("Error flushing csv writer with 'writer.Error()' in 'WriteHeadersToCsvFile'. Error: %v", err)
	}

	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// Validate headers of csv file
// /////////
func ValidateCSVHeaders(csvFile *os.File, expectedHeaders []string) error {
	reader := csv.NewReader(csvFile)
	headerRow, err := reader.Read()
	if err != nil {
		return fmt.Errorf("Error reading CSV header row in 'ValidateCSVHeaders()': %v", err)
	}

	if !reflect.DeepEqual(headerRow, expectedHeaders) {
		return fmt.Errorf("CSV headers do not match expected headers. Expected: %v, Actual: %v", expectedHeaders, headerRow)
	}

	return nil
}
