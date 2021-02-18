package main

import (
	_ "embed"
	"fmt"
	_ "net/http/pprof"
	"unsafe"
)

/*-------------------INPUT-------------------*/
//var reader = bufio.new(os.Stdin)
//var writer = bufio.NewWriter(os.Stdout)

//var writer = bufio.NewWriter(os.Stdout)
//var reader = bufio.NewScanner(os.Stdin)

//go:embed n2.txt
var s string

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	str := "12313"
	ss1 := *(*[]byte)(unsafe.Pointer(&str))
	fmt.Println(ss1)

	ss2 := *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			asdad int
			//Cap int
		}{str, len(str)},
	))
	fmt.Println(ss2)
}
