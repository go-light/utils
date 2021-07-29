package sqlutils

import (
	"fmt"
	"regexp"
	"strings"
)

func Joint(query string, fieldValue []string) string {
	r := regexp.MustCompile("(?i)from")
	query = r.ReplaceAllString(query, "FROM")
	query = strings.Trim(query, ";")

	newQuery := ""
	for _, t := range fieldValue {
		newQuery += strings.Replace(query, " FROM", fmt.Sprintf(",%s as `t` FROM", t), 1) + " UNION ALL "
	}

	newQuery = strings.Trim(newQuery, " UNION ALL ") + ";"
	return newQuery
}


