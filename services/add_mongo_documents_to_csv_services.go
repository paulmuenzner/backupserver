package services

import (
	handleBson "backupserver/utils/bson"
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WriteBSONToCSV(bsonData []bson.M, csvFilePath string) error {
	// Create or open the CSV file
	file, err := os.OpenFile(csvFilePath, os.O_RDWR, 0644)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("Permission denied to open csv file: %s in 'WriteBSONToCSV'. Error: %v", csvFilePath, err)
		} else {
			return fmt.Errorf("Error opening csv file path: %s in 'WriteBSONToCSV'. Error: %v", csvFilePath, err)
		}
	}
	defer file.Close()

	// Initialize CSV writer
	writer := csv.NewWriter(file)

	// Read existing CSV header if it exists
	var existingHeader []string
	existingData, err := csv.NewReader(file).Read()
	if err == nil {
		existingHeader = existingData
	}

	// Write header if not already present
	if len(existingHeader) == 0 {
		var header []string
		for key := range handleBson.FlattenBSON(bsonData[0], "", ".") {
			header = append(header, key)
		}
		if err := writer.Write(header); err != nil {
			return fmt.Errorf("Error writing header to csv file with 'writer.Write()' in 'WriteBSONToCSV'. File name: %s. Error: %v\n", csvFilePath, err)
		}
		existingHeader = header
	}

	// Create a channel to communicate between goroutines
	rowChannel := make(chan []string, len(bsonData))

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch a goroutine for each BSON entry
	for _, entry := range bsonData {
		wg.Add(1)
		go func(entry bson.M) {
			defer wg.Done()

			// Flatten BSON entry
			flatEntry := handleBson.FlattenBSON(entry, "", ".")

			// Prepare row data based on existing header
			var row []string
			for _, headerKey := range existingHeader {
				if value, exists := flatEntry[headerKey]; exists {
					if oid, ok := value.(primitive.ObjectID); ok {
						row = append(row, oid.Hex())
					} else {
						row = append(row, fmt.Sprintf("%v", value))
					}
				} else {
					row = append(row, "")
				}
			}

			// Send the row data to the channel
			rowChannel <- row
		}(entry)
	}

	// Close the channel after all goroutines have finished
	go func() {
		wg.Wait()
		close(rowChannel)
	}()

	// Receive rows from the channel and write to CSV
	for row := range rowChannel {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("Error writing received rows from channel with 'writer.Write()' in 'WriteBSONToCSV'. File name: %s. Error: %v\n", csvFilePath, err)
		}
	}

	// Flush and close CSV writer
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("Error flushing and closing csv writer with 'writer.Flush()' in 'WriteBSONToCSV'. File name: %s. Error: %v\n", csvFilePath, err)
	}

	return nil
}
