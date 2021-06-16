package urlutils

import (
	"fmt"
	"net/url"
	"strings"
)

func Joint(schemeHost string, path string, urlValuesEncode string) string {
	return strings.TrimRight(schemeHost, "/") + "/" +
		strings.TrimLeft(path, "/") + "?" +
		strings.TrimLeft(urlValuesEncode, "?")
}

func SetRawQuery(oldUrl string, key string, value string) (string, error) {
	u, err := url.Parse(oldUrl)
	if err != nil {
		return "", err
	}

	urlValues := u.Query()
	urlValues.Set(key, value)

	newUrl := ""
	if u.User.String() == "" {
		newUrl = fmt.Sprintf("%s://%s%s?%s#%s", u.Scheme, u.Host, u.Path, urlValues.Encode(), u.Fragment)

	} else {
		newUrl = fmt.Sprintf("%s://%s@%s%s?%s#%s", u.Scheme, u.User, u.Host, u.Path, urlValues.Encode(), u.Fragment)

	}

	return newUrl, nil
}
