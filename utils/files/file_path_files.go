package files

import "github.com/paulmuenzner/backupserver/config"

// Get entire file path depending on timestamp, file name and name of backup folder
func GetFilePath(fileName string, timeStamp string) (filePath string) {
	fileEnding := ".csv"
	pathSubFolder := GetSubFolder(timeStamp)
	filePath = pathSubFolder + "/" + fileName + fileEnding

	return filePath

}

// Create name of subfolder path
func GetSubFolder(timeStamp string) (pathSubFolder string) {
	folder := config.FolderNameBackup
	subFolder := "/" + timeStamp
	pathSubFolder = folder + subFolder

	return pathSubFolder

}
