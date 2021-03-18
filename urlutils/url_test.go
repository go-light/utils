package urlutils

import (
	"fmt"
	url4 "net/url"
	"testing"
)

func TestJoint(t *testing.T) {
	a := url4.Values{}
	a.Encode()
	url := Joint("http://m.baidu.com/", "/a/b", "?a=b")
	fmt.Println(url)

	url1 := Joint("http://m.baidu.com/", "a/b", "?as")
	fmt.Println(url1)

	url2 := Joint("http://m.baidu.com", "/a/b", "a=b")
	fmt.Println(url2)

	url3 := Joint("http://m.baidu.com", "a/b", "c=d")
	fmt.Println(url3)
}
