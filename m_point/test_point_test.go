package m_point

import (
	"fmt"
	"testing"
)

func TestPoint_(t *testing.T) {
	a := 1
	p := &a
	*p++
	fmt.Println(a, p)
}

func TestPointArr_(t *testing.T) {
	TestPointArr()
}
