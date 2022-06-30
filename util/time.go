package util

import (
	"time"
)

//current week first day with location utc
func WeekMondayTimestamp(now time.Time) int64 {
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Unix()
}

func TimeStamp(t string) int64 {
	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	return stamp.Unix()
}

func FromUnixtime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
