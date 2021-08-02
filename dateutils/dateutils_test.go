package dateutils

import (
	"fmt"
	"testing"
	"time"
)

func TestParseOut(t *testing.T) {
	d := ParseOut(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 31, 23, 59, 59, 0, time.UTC), TypeDay)
	fmt.Println(d.List)

	//fmt.Println(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Format(timeutils.Template))
	//fmt.Println(time.Date(2021, 8, 7, 23, 59, 59, 0, time.UTC).Format(timeutils.Template))

	w := ParseOut(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 8, 23, 59, 59, 0, time.UTC), TypeWeek)
	fmt.Println(w.List, w.Describe, w.StartDate, w.EndDate)

	m := ParseOut(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 31, 23, 59, 59, 0, time.UTC), TypeMonth)
	fmt.Println(m.List)
}
