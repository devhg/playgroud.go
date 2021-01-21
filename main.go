package main

import (
	"fmt"
	"math"
	"strings"
)

/*-------------------INPUT-------------------*/
//var reader = bufio.new(os.Stdin)
//var writer = bufio.NewWriter(os.Stdout)

//var writer = bufio.NewWriter(os.Stdout)
//var reader = bufio.NewScanner(os.Stdin)

func foo() int {
	var a int64 = 91283472332
	if a > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(a)
}

func main() {
	//var a int64 = 10
	//var b byte = '9'
	//a += a*10 + int64(b - '0')
	strings.TrimSpace()
	fmt.Print(foo())
}
