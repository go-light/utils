package sqlutils

import (
	"fmt"
	"regexp"
	"strings"
)

func Joint(query string, fieldValue []string) string {
	r := regexp.MustCompile("(?i)from")
	query = r.ReplaceAllString(query, "FROM")

	r1 := regexp.MustCompile("(?i);")
	query = r1.ReplaceAllString(query, "")

	newQuery := ""
	for _, t := range fieldValue {
		newQuery += strings.Replace(query, " FROM", fmt.Sprintf(",'%s' as `x_arg_date` FROM", t), 1) + " UNION ALL "
	}

	newQuery = strings.Trim(newQuery, " UNION ALL ") + ";"
	return newQuery
}
