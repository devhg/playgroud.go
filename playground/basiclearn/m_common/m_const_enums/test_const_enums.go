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

	const (
		aa = iota   // iota = 0
		bb          // iota = 1
		cc          // iota = 2
		dd = "a21a" //独立值 iota += 1
		ee          //独立值 iota += 1
		ff = iota   // 返回iota += 1
		gg          // iota += 1
	)
	//0 1 2 a21a a21a 5 6
	//通过该案例可以明显看到iota遇到主动赋值的条目时，并不会终止累加，而是会继续隐式增加iota的值。
	fmt.Println(aa, bb, cc, dd, ee, ff, gg)
}
