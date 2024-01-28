package config

const (
	DeleteLogsAfterDays       = 5
	NameDatabase              = "PlantDB"
	FolderNameBackup          = "plantdb"
	FileNameMetaData          = "meta_data.csv"
	IntervalBackup            = "@every 10s" // cron-like syntax format to define a recurring schedule. Alternative examples: "@every 30m", "@every 6h", "@every 1d"
	MongoURIEnv               = "MONGO_URI"
	MaxFileSizeInBytes  int64 = 2 /* <<< Size in GB*/ * 1024 * 1024 * 1024 // Maximum permitted backup csv file size. Consider max upload size permitted by AWS S3 => README.md
	// Email
	SendEmailNotifications   = false // If false, no email notifications at all (error & success)
	EmailProviderUserNameEnv = "EMAIL_PROVIDER_USERNAME"
	EmailProviderPasswordEnv = "EMAIL_PROVIDER_PASSWORD"
	EmailProviderSmtpPortEnv = "EMAIL_PROVIDER_SMTP_PORT"
	EmailProviderHostEnv     = "EMAIL_PROVIDER_HOST"
	EmailAddressSenderEnv    = "EMAIL_ADDRESS_SENDER_BACKUP"
	EmailAddressReceiverEnv  = "EMAIL_ADDRESS_RECEIVER_BACKUP"
	// Circular buffer S3 settings
	IsCircularBufferActivatedS3 = true // If false, all created backups will be stored and not deleted by this program - no circular buffer mechanism
	MaxBackupsS3                = 5    // Circular buffer deletes backups older than latest number of MaxBackupsS3 in S3
	// Circular buffer local settings
	UseLocalBackupStorage            = true // If false, backups are only stored on aws
	IsCircularBufferActivatedLocally = true // If false, all created backups will be stored and not deleted by this program - no circular buffer mechanism
	MaxBackupsLocally                = 5    // Circular buffer deletes backups older than latest number of MaxBackupsLocally locally if above 'LocalBackupStorage' is true
	// AWS S3 Production config .env variable names
	S3BucketEnv    = "BUCKET_NAME"
	S3RegionEnv    = "AWS_REGION"
	S3AccessKeyEnv = "AWS_ACCESS_KEY_ID"
	S3SecretKeyEnv = "AWS_SECRET_ACCESS_KEY"
)
