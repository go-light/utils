package sqlutils

import (
	"fmt"
	"testing"
)

func TestJoint(t *testing.T) {
		fieldValues := []string{"2012", "2013", "2014"}
		query := "select * from tables WHERE start_time =? and end_time =? group by m ;"
		fmt.Println(Joint(query, fieldValues))
}
