package algorithm

import (
	"fmt"
	"testing"
)

func TestMains(t *testing.T) {
	slice := make([]int, 0)
	slice = append(slice, 1)

	p := slice[0:]
	fmt.Println(p)
}
