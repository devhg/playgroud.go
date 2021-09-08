package datastruct

import (
	"fmt"
)

func TestSlice() {
	/*
		arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
		printSlice(arr[:]) //数组转切片并打印

		s := arr[2:6] // 切片
		fmt.Println(reflect.TypeOf(s))
		// 下面都是切片
		fmt.Println("arr[2:6]", arr[2:6])
		fmt.Println("arr[2:]", arr[2:])
		fmt.Println("arr[:6]", arr[:6])
		fmt.Println("arr[:]", arr[:])

		fmt.Println("s", s)
		updateSlice(s)
		fmt.Println("after update slice")
		fmt.Println("s", s)
		fmt.Println(arr) // 改变arr的切片 arr也会被改变
	*/
	extendSlice()
	createSlice()
	copyDelSlice()
}

// []int是切片类型
func updateSlice(s []int) {
	s[0] = 1000
}

// 打印切片
func printSlice(s []int) {
	// for i, v := range s {
	// 	fmt.Println(i, " ", v)
	// }
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

// 切片的扩展
// 添加元素的时候如果超过cap，系统会重新分配更大的底层数组，把内容拷贝过去
// 原来的数组若不在使用，将会被gc
// 由于值传递的关系，必须接受append的返回值
func extendSlice() {
	fmt.Println("Extending slice")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // {2, 3, 4, 5} 6 7
	s2 := s1[3:5]  // {5, 6} 7
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1)) // s1=[2 3 4 5], len(s1)=4, cap(s1)=6

	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2)) // s2=[5 6], len(s2)=2, cap(s2)=3

	// 这里没有越界的原因看slice底层原理

	s3 := append(s2, 10) // 因为原来的slice 里存在7 所以10会换掉7
	s4 := append(s3, 11) // 这里超过了原来arr的容量，就会创建一个新的arrnew，把原来的拷贝过去
	s5 := append(s4, 12)
	fmt.Println("s3 s4 s5= ", s3, s4, s5)
	fmt.Println("arr= ", arr)
}

// 切片的创建
func createSlice() {
	fmt.Println("Creating slice")
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, i)
		// 以2的指数形式扩增
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	}
	fmt.Printf("s=%v \n", s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 16, 32)
	printSlice(s3)
}

// 切片的拷贝和删除
func copyDelSlice() {
	fmt.Println("Copying slice")
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	s2 := make([]int, 16)
	printSlice(s2)

	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) // 由于append只能接收可边长参数  s2[4:]...是将切片转成可边长参数
	printSlice(s2)                 // 8被删除

	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
