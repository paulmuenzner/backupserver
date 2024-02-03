package files

import (
	"os"
	"testing"
)

func TestDeleteLocalFolder(t *testing.T) {
	// Create a temporary folder for testing
	testFolder := "test_folder"
	err := os.Mkdir(testFolder, 0750)
	if err != nil {
		t.Fatalf("Failed to create test folder: %v", err)
	}
	defer os.RemoveAll(testFolder)

	tests := []struct {
		name     string
		folder   string
		expected error
	}{
		{"DeleteExistingFolder", testFolder, nil},
		{"DeleteNonExistentFolder", "nonexistent_folder", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteLocalFolder(tt.folder)

			if (err != nil) != (tt.expected != nil) {
				t.Errorf("DeleteLocalFolder() error = %v, expected %v", err, tt.expected)
				return
			}

			if err != nil && err.Error() != tt.expected.Error() {
				t.Errorf("DeleteLocalFolder() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}
