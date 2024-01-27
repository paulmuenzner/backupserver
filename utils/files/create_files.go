package files

// // Create empty file of any type without header using full file path, for example "backupserver/folder/subfolder/example.csv"
// func CreateEmptyFileAndClose(filePath string) (err error) {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	return nil
// }

// // Create empty file of any type without header using full file path, for example "backupserver/folder/subfolder/example.csv"
// func CreateEmptyFileAndKeepOpen(filePath string) (file *os.File, err error) {
// 	file, err = os.Create(filePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return file, nil
// }
