package m_panic_defer_recover

import (
	"fmt"
)

//panic recover
//停止当前函数执行
//一直向上返回，执行每一层defer
//如果没有遇见recover，程序退出

//re
func coverPanic() {
	message := recover()
	switch v := message.(type) {
	case string:
		fmt.Println("string message: ", message, v)
	case error:
		fmt.Println("error message: ", message, v)
	case int:
		fmt.Println("int message: ", message, v)
	default:
		fmt.Println("unknown message: ", message, v)
		panic(message) // 处理不了 重新panic
	}
}

func TestPanic_() {
	defer coverPanic()
	//panic(errors.New("I am m_panic_defer_recover"))

	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	panic("123")
}

func TestDefer_() {
	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println(3)
	panic("error")
	fmt.Println(4)
	// 3 2 1
	// 多个defer的时候按照压栈、出栈顺序进行
}
