package urlutils

import (
	"fmt"
	"testing"
)

func TestJoint(t *testing.T) {
	url := Joint("http://m.baidu.com/", "/a/b", "?a=b")
	fmt.Println(url)

	url1 := Joint("http://m.baidu.com/", "a/b", "?as")
	fmt.Println(url1)

	url2 := Joint("http://m.baidu.com", "/a/b", "a=b")
	fmt.Println(url2)

	url3 := Joint("http://m.baidu.com", "a/b", "c=d")
	fmt.Println(url3)
}

func TestAppendRawQuery(t *testing.T)  {
	u, err := SetRawQuery("postgres://user:pass@host.com:5432/path?k=v#f", "k1", "v1")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(u)

	u1, err1 := SetRawQuery("https://host.com:5432/path?k=v#f", "k1", "v1")
	if err1 != nil {
		t.Error(err1)
		return
	}

	fmt.Println(u1)

	u1, err1 = SetRawQuery(".com:5432/path?k=v#f", "k1", "v1")
	if err1 != nil {
		t.Error(err1)
		return
	}

	fmt.Println(u1)
}
