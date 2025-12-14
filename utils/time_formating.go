package utils

import "time"

func GetDateFormat(date time.Time) string {
	return date.Format("2006-01-02")
}
