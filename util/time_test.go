package util

import (
	"testing"
	"time"
)

func init() {
	//time.Local = time.UTC
}

func TestWeekMondayTimestamp(t *testing.T) {
	monday, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-06-27 00:00:00", time.Local)
	var nows [7]time.Time
	nows[0] = FromUnixtime(TimeStamp("2022-06-27 15:04:05"))
	nows[1] = FromUnixtime(TimeStamp("2022-06-28 15:04:05"))
	nows[2] = FromUnixtime(TimeStamp("2022-06-29 15:04:05"))
	nows[3] = FromUnixtime(TimeStamp("2022-06-30 15:04:05"))
	nows[4] = FromUnixtime(TimeStamp("2022-07-01 15:04:05"))
	nows[5] = FromUnixtime(TimeStamp("2022-07-02 15:04:05"))
	nows[6] = FromUnixtime(TimeStamp("2022-07-03 15:04:05"))

	for i := 0; i < len(nows); i++ {
		if WeekMondayTimestamp(nows[i]) != monday.Unix() {
			t.Fatal("not monday", nows[i], monday.Unix())
		}
	}
}
