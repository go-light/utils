package sqlutils

import (
	"fmt"
	"testing"
)

func TestJoint(t *testing.T) {
		fieldValues := []string{"2012", "2013", "2014"}
		query := "Select sum(single_question_num), yearweek(stat_date) from (\nSELECT *, yearweek(stat_date) as w FROM warehouse.dws_pingce_core where stat_date >= \"2021-07-16 00:00:00\" and stat_date <= \"2021-07-22 00:00:00\") t group by w;\n;"
		fmt.Println(Joint(query, fieldValues))
}
