package m_functional

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	f := add()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

func Test_adder(t *testing.T) {
	f := adder(0)

	for i := 0; i < 10; i++ {
		var s int
		s, f = f(i)
		fmt.Println(s)
	}
}

func Test_fibonacci(t *testing.T) {
	fib := fibonacci()
	printFileContents(fib)
}
