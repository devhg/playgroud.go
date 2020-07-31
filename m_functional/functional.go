package m_functional

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//函数式编程，闭包

func add() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

//更标准的函数是编程
type iAdd func(int) (int, iAdd)

func adder(base int) iAdd {
	return func(v int) (int, iAdd) {
		return base + v, adder(base + v)
	}
}

/*=========================================*/
//闭包、斐波那契数列、和reader接口示例
type intFib func() int

func (i intFib) Read(p []byte) (n int, err error) {
	next := i()

	if next > 10000 {
		return 0, io.EOF
	}
	return strings.NewReader(fmt.Sprintf("%d\n", next)).Read(p)
}

func fibonacci() intFib {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
