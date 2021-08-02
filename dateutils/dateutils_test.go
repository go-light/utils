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

func TestParseOutWeek(t *testing.T) {
	w := ParseOut(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 8, 23, 59, 59, 0, time.UTC), TypeWeek, WithWeekMode(2))
	fmt.Println(w.List, w.Describe, w.StartDate, w.EndDate)

	w1 := ParseOut(time.Date(2021, 8, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 8, 23, 59, 59, 0, time.UTC), TypeWeek, WithWeekMode(2))
	fmt.Println(w1.List, w1.Describe, w1.StartDate, w1.EndDate)

	w2 := ParseOut(time.Date(2021, 8, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 10, 23, 59, 59, 0, time.UTC), TypeWeek, WithWeekMode(6))
	fmt.Println(w2.List, w2.Describe, w2.StartDate, w2.EndDate)

	w3 := ParseOut(time.Date(2021, 8, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 9, 23, 59, 59, 0, time.UTC), TypeWeek, WithWeekModeByLastDayOfWeek(time.Now()))
	fmt.Println(w3.List, w3.Describe, w3.StartDate, w3.EndDate)

}
