package m_point

import "fmt"

type t struct {
	a string
}

func TestPoint() {
	var a int = 999
	var ap *int
	ap = &a
	var app **int = &ap

	fmt.Printf("a地址  %x\n", &a)
	fmt.Printf("ap值   %x\n", ap)

	fmt.Printf("ap取值运算 %d\n", *ap)
	fmt.Printf("app取值运算 %d\n", **app)

	var test = new(t)
	fmt.Println(test)
}

func TestPointArr() {
	// 指针数组
	a, b := 1, 2
	arr := [...]*int{&a, &b}
	fmt.Println("指针数组 : ", arr)
	for i := range arr {
		fmt.Println(*arr[i])
	}

	// 数组指针
	parr := &[...]int{1, 2}
	fmt.Println("数组指针 : ", parr)
	for i := range *parr {
		fmt.Println((*parr)[i])
	}

}
