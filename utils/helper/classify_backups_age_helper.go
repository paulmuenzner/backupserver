package helper

import "sort"

// Determine names of backup folders to delete depending on max. number of permitted backups
func ClassifyBackupsByAge(backupDates []string, maxNumberBackups int) (outdatedBackups []string, latestBackups []string) {
	// Sort the slice in descending order
	sort.Slice(backupDates, func(i, j int) bool {
		return backupDates[i] > backupDates[j]
	})

	// Create list of newest and oldest backups
	outdatedBackups = backupDates[maxNumberBackups:]
	latestBackups = backupDates[:maxNumberBackups]

	return outdatedBackups, latestBackups
}
