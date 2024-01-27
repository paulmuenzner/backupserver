package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Count and list folder paths locally inside specified folder path
func CountAndListCSVFiles(folderPath string) (int, []string, error) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return 0, nil, err
	}

	csvFileCount := 0
	csvFilePaths := make([]string, 0)

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
			csvFileCount++
			filePath := filepath.Join(folderPath, file.Name())
			csvFilePaths = append(csvFilePaths, filePath)
		}
	}

	return csvFileCount, csvFilePaths, nil
}

// List all folder names as string slice inside a given local path
func ListLocalFolderNames(folderPath string) ([]string, error) {
	// Validate if folder path really exists
	folderExists, err := LocalFolderExists(folderPath)
	if err != nil {
		return nil, fmt.Errorf("Error validating if folder exists in 'ListLocalFolderNames'. Folder path %s. Error: %v\n", folderPath, err)
	} else if folderExists == false {
		return nil, fmt.Errorf("Folder with path path %s not existing. Validated in 'ListLocalFolderNames' with 'LocalFolderExists()'.", folderPath)
	}

	var folderNames []string
	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Error walking the file tree inside folder %s. Error: %v\n", folderPath, err)
		}

		if info.IsDir() && path != folderPath {
			if _, err := os.Stat(path); err == nil { // Re-check existence before adding
				folderNames = append(folderNames, filepath.Base(path))
			} else {
				return fmt.Errorf("Folder not found in 'ListLocalFolderNames': %s\n", path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return folderNames, nil
}
