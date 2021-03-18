package timeutils

import "time"

const Template = "2006-01-02 15:04:05"

func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
