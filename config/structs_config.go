package config

import "time"

// ////////////////////////////////////////////////////////////////////////
// RowTypes represents a row of data in the CSV file for meta information
// /////////////////
type RowTypesMeta struct {
	CollectionName    string    //`csv:"collectionName"`
	FolderPath        string    //`csv:"folderPath"`
	FileName          string    //`csv:"fileName"`
	SizeInBytes       int64     //`csv:"sizeInBytes"`
	DatabaseName      string    //`csv:"databaseName"`
	DateStartBackup   time.Time //`csv:"dateStartBackup"`
	DateFinalizedFile time.Time //`csv:"dateFinalizedFile"`
}

// MetaHeaders represents the header names for the meta CSV file
var MetaFileHeaders = []string{
	"collection_name",
	"folder_path",
	"file_name",
	"size_in_bytes",
	"database_name",
	"date_start_entire_backup",
	"date_finalized_file",
}
