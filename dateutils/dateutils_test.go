package dateutils

import (
	"fmt"
	"testing"
	"time"
)

func TestParseOut(t *testing.T) {
	m := ParseOut(time.Now(), time.Now().Add(35*24*time.Hour), TypeWeek)
	fmt.Println(m)
}
