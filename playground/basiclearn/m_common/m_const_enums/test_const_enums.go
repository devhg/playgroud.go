package m_const_enums

import (
	"fmt"
	"math"
)

func Const() {
	const (
		filename = "123.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func Enum() {
	// go的枚举类型
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)

	// 自增值  _跳过
	const (
		cpp_ = iota
		_
		python_
		golang_
	)
	fmt.Println(cpp_, python_, golang_)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
