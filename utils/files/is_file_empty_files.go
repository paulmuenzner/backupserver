package files

import "os"

// Validate if file of any type is empty by using Go file descriptor as parameter type
func IsFileEmptyByDescriptor(file *os.File) (bool, error) {
	fileInfo, err := GetFileSizeByDescriptor(file)
	if err != nil {
		return false, err
	}

	return fileInfo == 0, nil
}

// Validate if file of any type is empty by using file path of type string
func IsFileEmptyByFilePath(filePath string) (bool, error) {
	fileInfo, err := GetFileSizeByPath(filePath)
	if err != nil {
		return false, err
	}

	return fileInfo == 0, nil
}
