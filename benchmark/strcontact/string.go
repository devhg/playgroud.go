package strcontact

import (
	"bytes"
	"fmt"
)

var a string = "hello "
var b string = "world"

// 字符串相加
func StrContactByAdd() string {
	return a + b
}

// Sprintf
func StrContactByfmSpf() string {
	return fmt.Sprintf("%s%s", a, b)
}

//bytes.Buffer
func StrContactByBuffer() string {
	var buffer bytes.Buffer
	buffer.WriteString(a)
	buffer.WriteString(b)
	return buffer.String()
}

// slice append then to string
func StrContactByAppend() string {
	return string(append([]byte(a), []byte(b)...))
}
