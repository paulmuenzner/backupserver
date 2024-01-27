package date

import "time"

func TimeStampSlug(timeStamp time.Time) string {
	date := timeStamp.Format("2006-01-02")
	hours := timeStamp.Format("15")
	minutes := timeStamp.Format("04")
	seconds := timeStamp.Format("05")
	return date + "_" + hours + "-" + minutes + "-" + seconds
}

func TimeStamp() time.Time {
	return time.Now()
}
