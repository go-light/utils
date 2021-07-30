package dateutils

import (
	"fmt"
	"github.com/go-light/utils/timeutils"
	"time"
)

type ParseOutType string

const (
	TypeMonth ParseOutType = "month"
	TypeWeek  ParseOutType = "week"
	TypeDay   ParseOutType = "day"
)

type DurationDate struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Result struct {
	StartDate string         `json:"start_date"`
	EndDate   string         `json:"end_date"`
	List      []string       `json:"list"` // header
	Data      []string       `json:"data"` // data
	Describe  []DurationDate `json:"describe"`
}

func ParseOut(startTime time.Time, endTime time.Time, parseOutType ParseOutType, opts ...Option) (ret *Result) {
	options := &options{
		weekMode: 0,
	}

	for _, opt := range opts {
		opt(options)
	}

	ret = &Result{}
	switch parseOutType {
	case TypeMonth:
		ret = parseOutMonth(startTime, endTime)
	case TypeWeek:
		ret = parseOutWeek(startTime, endTime)
	case TypeDay:
		ret = parseOutDay(startTime, endTime)
	}
	return ret
}

const (
	Sunday    = "Sunday"
	Monday    = "Monday"
	Tuesday   = "Tuesday"
	Wednesday = "Wednesday"
	Thursday  = "Thursday"
	Friday    = "Friday"
	Saturday  = "Saturday"
)

var weekdayToNumber = map[string]int{
	Sunday:    0,
	Monday:    1,
	Tuesday:   2,
	Wednesday: 3,
	Thursday:  4,
	Friday:    5,
	Saturday:  6,
}

func parseOutWeek(startTime time.Time, endTime time.Time) (ret *Result) {
	ret = &Result{}
	list := make([]string, 0, 20)
	describe := make([]DurationDate, 0, 20)

	start := startTime.AddDate(0, 0, -weekdayToNumber[startTime.Weekday().String()])

	timeStr := start.Format("2006-01-02")
	start, _ = time.Parse("2006-01-02", timeStr)

	end := endTime.AddDate(0, 0, -weekdayToNumber[startTime.Weekday().String()]+6)

	ret.StartDate = start.Format(timeutils.Template)
	ret.EndDate = end.Format(timeutils.Template)

	tmpStart := start
	for {
		y, w := tmpStart.ISOWeek()
		tmpFormat := fmt.Sprintf("%d%02d", y, w)
		tmpEnd := tmpStart.AddDate(0, 0, +6).Add(24*time.Hour - 1)

		dd := DurationDate{
			StartDate: tmpStart.Format(timeutils.Template),
			EndDate:   tmpEnd.Format(timeutils.Template),
		}

		list = append(list, tmpFormat)
		describe = append(describe, dd)

		tmpStart = tmpStart.AddDate(0, 0, 7)
		if tmpStart.Unix() >= end.Unix() {
			break
		}
	}

	ret.List = list
	ret.Describe = describe

	return
}

func parseOutDay(startTime time.Time, endTime time.Time) (ret *Result) {
	ret = &Result{}
	list := make([]string, 0, 60)
	describe := make([]DurationDate, 0, 60)

	start := startTime
	timeStr := start.Format("2006-01-02")
	start, _ = time.Parse("2006-01-02", timeStr)

	end := endTime
	endStr := end.Format("2006-01-02")
	end, _ = time.Parse("2006-01-02", endStr)

	ret.StartDate = start.Format(timeutils.Template)
	ret.EndDate = end.Format(timeutils.Template)

	tmpStart := start
	for {
		tmpFormat := tmpStart.Format("2006-01-02")
		tmpEnd := tmpStart.AddDate(0, 0, +1).Add(-1)

		dd := DurationDate{
			StartDate: tmpStart.Format(timeutils.Template),
			EndDate:   tmpEnd.Format(timeutils.Template),
		}

		list = append(list, tmpFormat)
		describe = append(describe, dd)

		tmpStart = tmpStart.AddDate(0, 0, 1)
		if tmpStart.Unix() >= end.Unix() {
			break
		}
	}

	ret.List = list
	ret.Describe = describe

	return
}

func parseOutMonth(startTime time.Time, endTime time.Time) (ret *Result) {
	ret = &Result{}
	list := make([]string, 0, 20)
	describe := make([]DurationDate, 0, 20)

	year, month, _ := startTime.Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	year1, month1, _ := endTime.Date()
	end := time.Date(year1, month1, 1, 23, 59, 59, 0, time.UTC).AddDate(0, 1, -1)

	ret.StartDate = start.Format(timeutils.Template)
	ret.EndDate = end.Format(timeutils.Template)

	tmpStart := start
	for {
		tmpMonth := tmpStart.Format("2006-01")
		tmpEnd := tmpStart.AddDate(0, 1, -1).Add(24*time.Hour - 1)

		dd := DurationDate{
			StartDate: tmpStart.Format(timeutils.Template),
			EndDate:   tmpEnd.Format(timeutils.Template),
		}

		list = append(list, tmpMonth)
		describe = append(describe, dd)

		tmpStart = tmpStart.AddDate(0, 1, 0)
		if tmpStart.Unix() >= end.Unix() {
			break
		}
	}

	ret.List = list
	ret.Describe = describe

	return
}
