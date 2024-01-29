package files

import "os"

/////////////////////////////////////////////////////////////
// Check if folder exists
///////////
func LocalFolderExists(folderPath string) (bool, error) {
	info, err := os.Stat(folderPath)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/////////////////////////////////////////////////////////////
// Check if local file exists
///////////
func LocalFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil // Path exists (file or folder)
	}
	if os.IsNotExist(err) {
		return false, nil // Path does not exist
	}
	return false, err // Other error
}
