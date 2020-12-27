package stack

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := Stack{}
	s.Push(121)
	fmt.Println(s.Pop())
}
