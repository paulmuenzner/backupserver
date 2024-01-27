package files

import (
	"os"
)

// Get file size in bytes of any file type by using Go file descriptor as parameter type
func GetFileSizeByPath(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err // Prepare
	}
	return fileInfo.Size(), nil
}

// Get file size in bytes of any file type by using Go file descriptor as parameter type
func GetFileSizeByDescriptor(file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err // Prepare
	}
	return fileInfo.Size(), nil
}
