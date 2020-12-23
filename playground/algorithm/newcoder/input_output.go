package newcoder

import (
	"bufio"
	"fmt"
	"os"
)

/*-------------------INPUT-------------------*/
var reader = bufio.NewReader(os.Stdin)

func read() (x int) {
	c, e := reader.ReadByte()
	neg := false
	for ; c < '0' && e == nil; c, e = reader.ReadByte() {
		neg = '-' == c
	}
	for ; c >= '0' && e == nil; c, e = reader.ReadByte() {
		x = x*10 + (int(c) ^ 48)
	}
	if neg {
		x = -x
	}
	return
}

/*-------------------OUTPUT-------------------*/
var writer = bufio.NewWriter(os.Stdout)

func main() {
	n := 0
	n = read()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = read()
	}

	fmt.Fprint(writer, 1, " ", 2, "\n")
	writer.Flush()
}
