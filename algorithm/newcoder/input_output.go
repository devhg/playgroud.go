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

func read2() {
	reader := bufio.NewReader(os.Stdin)
	cnt := 0
	for b, err := reader.ReadByte(); b >= ' ' && err == nil; b, err = reader.ReadByte() {
		if b == '(' {
			cnt++
		} else if b == ')' && cnt > 0 {
			cnt--
		} else {
			fmt.Println("NO")
			return
		}
	}
	if cnt == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
