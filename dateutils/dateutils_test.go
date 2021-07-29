package dateutils

import (
	"fmt"
	"testing"
	"time"
)

func TestParseOut(t *testing.T) {
	m := ParseOutMonth(time.Now(), time.Now().Add(35*24*time.Hour))
	fmt.Println(m)
}
