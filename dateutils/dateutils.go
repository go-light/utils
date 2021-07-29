package dateutils

import (
	"github.com/go-light/utils/timeutils"
	"time"
)

type Week struct {
}

type Day struct {
}

type DurationDate struct {
	StartDate string`json:"start_date"`
	EndDate   string `json:"end_date"`
}

type MonthArgs struct {
	StartDate string                  `json:"start_date"`
	EndDate   string                  `json:"end_date"`
	List      []string                `json:"list"` // header
	Data      []string                `json:"data"` // data
	Describe  []DurationDate `json:"describe"`
}

func ParseOutMonth(startTime time.Time, endTime time.Time) *MonthArgs {
	m := MonthArgs{}
	list := make([]string, 0, 20)
	describe := make([]DurationDate, 0, 20)

	year, month, _ := startTime.Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	year1, month1, _ := endTime.Date()
	end := time.Date(year1, month1, 1, 23, 59, 59, 0, time.UTC).AddDate(0, 1, -1)

	m.StartDate = start.Format(timeutils.Template)
	m.EndDate = end.Format(timeutils.Template)

	tmpStart := start
	for {
		tmpMonth := tmpStart.Format("2006-01")
		tmpEnd := tmpStart.AddDate(0, 1, -1).Add(23*time.Hour+59+time.Minute+59*time.Second)

		dd := DurationDate{
			StartDate: tmpStart.Format(timeutils.Template),
			EndDate: tmpEnd.Format(timeutils.Template),
		}

		list = append(list, tmpMonth)
		describe = append(describe, dd)

		tmpStart = tmpStart.AddDate(0, 1, 0)
		if tmpStart.Unix() >= end.Unix() {
			break
		}
	}

	m.List = list
	m.Describe = describe

	return &m
}
