package algorithm

import (
	"fmt"
	"testing"
)

func TestMains(t *testing.T) {
	//Main()
	a := 1
	test(&a)
	fmt.Println(a)
}

func test(a *int) {
	*a = 10
}
