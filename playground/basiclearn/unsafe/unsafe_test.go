package unsafe_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	stringHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&stringHeader))
}

func TestMain1(t *testing.T) {
	str := "hello 12313!"
	bytes := string2bytes(str)
	fmt.Println(bytes)
	strin := bytes2string(bytes)
	fmt.Println(strin)
}
