package services

import (
	"fmt"

	"github.com/paulmuenzner/backupserver/config"
	files "github.com/paulmuenzner/backupserver/utils/files"
	"github.com/paulmuenzner/backupserver/utils/helper"
	logger "github.com/paulmuenzner/backupserver/utils/logs"
)

func ManageBackupsLocally() error {
	localBackupFolder := config.FolderNameBackup + "/"
	useLocalBackupStorage := config.UseLocalBackupStorage
	if !useLocalBackupStorage {
		//////////////////////////////////////////////////////////////////////////////
		// Delete all locally cashed backups (saved locally used for S3 uploads)
		/////////////////////
		err := files.DeleteLocalFolder(localBackupFolder)
		if err != nil {
			return fmt.Errorf("Error in 'ManageBackupsLocally' utilizing 'DeleteLocalFolder()'. Error: %v", err)
		}
	} else {
		///////////////////////////////////////////////////////////////////////////////
		// Delete old local backups depending on if local circular buffer is activated
		/////////////////////
		isCircularBufferActivatedLocally := config.IsCircularBufferActivatedLocally
		if isCircularBufferActivatedLocally {
			//////////////////////////////////////////////////
			// Validate if more backups on local storage than permitted (accoding to config.MaxBackupsLocally)
			listLocalBackupFolders, err := files.ListLocalFolderNames(localBackupFolder)
			if err != nil {
				return fmt.Errorf("Error in 'ManageBackupsLocally' utilizing 'ListLocalFolderNames()'. Error: %v", err)
			}
			maxBackupsLocally := config.MaxBackupsLocally
			moreBackupsThanPermitted := len(listLocalBackupFolders) > maxBackupsLocally

			// If more local backups in backup folder than permitted, delete outdated ones
			if moreBackupsThanPermitted {
				// Get outdated backup folder names to delete
				outdatedBackups, _ := helper.ClassifyBackupsByAge(listLocalBackupFolders, maxBackupsLocally)

				// Delete outdated backups
				for _, outdatedBackupFolder := range outdatedBackups {
					path := localBackupFolder + outdatedBackupFolder
					err := files.DeleteLocalFolder(path)
					if err != nil {
						return fmt.Errorf("Error in 'ManageBackupsLocally' using 'DeleteLocalFolder(path)': %v", err)
					} else {
						logger.GetLogger().Infof("Delete backup folder %s", outdatedBackupFolder)
					}
				}
			}

		}
	}
	return nil
}
