package uaparser

import (
	"regexp"
	"sync"
)

// AppName/2.0.8(4997) OSType/1(Xiaomi M2102J2SC;11) NetType/WIFI Channel/xiaomi
const (
	OSTypeAndroid = "1"
	OSTypeIOS     = "2"
)

type Parser struct {
	AppVersion *AppVersion
	OS         *OS
}

func ParserUserAgent(userAgent string) *Parser {
	p := new(Parser)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		p.AppVersion = ParseAppVersion(userAgent)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		p.OS = ParserOs(userAgent)
	}()

	wg.Wait()

	return p
}

type OS struct {
	OsType string
}

func ParserOs(userAgent string) *OS {
	pat := `OSType/\d+`
	reg := regexp.MustCompile(pat)

	src := []byte(userAgent)
	template := []byte(`$0`)
	match := reg.FindSubmatchIndex(src)
	dst := reg.Expand(nil, template, src, match)

	pat1 := `\d+`
	reg1 := regexp.MustCompile(pat1)
	src1 := []byte(dst)
	template1 := []byte(`$0`)
	match1 := reg1.FindSubmatchIndex(src1)
	dst1 := reg1.Expand(nil, template1, src1, match1)

	osType := string(dst1)

	return &OS{
		OsType: osType,
	}
}

type Device struct {
}

type AppVersion struct {
	Version string
}

func ParseAppVersion(pat string, userAgent string) *AppVersion {
	//pat := `AppName/\d+\.\d+\.\d+`
	reg := regexp.MustCompile(pat)

	src := []byte(userAgent)
	template := []byte(`$0`)
	match := reg.FindSubmatchIndex(src)
	dst := reg.Expand(nil, template, src, match)

	pat1 := `\d+\.\d+\.\d+`
	reg1 := regexp.MustCompile(pat1)
	src1 := []byte(dst)
	template1 := []byte(`$0`)
	match1 := reg1.FindSubmatchIndex(src1)
	dst1 := reg1.Expand(nil, template1, src1, match1)

	version := string(dst1)

	return &AppVersion{
		Version: version,
	}
}
