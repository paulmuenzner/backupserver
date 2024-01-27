package files

import (
	"fmt"
	"os"
)

// Delete local folder entirely
func DeleteLocalFolder(folderToDelete string) error {
	err := os.RemoveAll(folderToDelete)
	if err != nil {
		return fmt.Errorf("Error in 'DeleteLocalFolder' deleting folder with 'os.RemoveAll()' of path: %s. Error: %v", folderToDelete, err)
	}

	return nil
}
