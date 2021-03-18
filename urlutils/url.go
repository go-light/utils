package urlutils

import "strings"

func Joint(schemeHost string, path string, urlValuesEncode string) string {
	return strings.TrimRight(schemeHost, "/") + "/" +
		strings.TrimLeft(path, "/") + "?" +
		strings.TrimLeft(urlValuesEncode, "?")
}
