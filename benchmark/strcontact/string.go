package strcontact

import (
	"bytes"
	"fmt"
	"strings"
)

var a []string

func init() {
	a = make([]string, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = "a "
	}
}

// 字符串相加
func StrContactByAdd() string {
	str := ""
	for _, s := range a {
		str += s
	}
	return str
}

// Sprintf
func StrContactByfmSpf() string {
	str := ""
	for _, s := range a {
		str = fmt.Sprintf("%s%s", str, s)
	}
	return str
}

// bytes.Buffer
func StrContactByBuffer() string {
	var buffer bytes.Buffer
	for _, s := range a {
		buffer.WriteString(s)
	}
	return buffer.String()
}

// strings.Builder
func StrContactByBuilder() string {
	var builder strings.Builder
	for _, s := range a {
		builder.WriteString(s)
	}
	return builder.String()
}

// slice append then to string
func StrContactByAppend() string {
	str := make([]byte, 0)
	for _, s := range a {
		str = append(str, []byte(s)...)
	}
	return string(str)
}
