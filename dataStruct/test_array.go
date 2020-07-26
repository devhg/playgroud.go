package dataStruct

import (
	"fmt"
	"reflect"
)

func TestArray() {
	var arr1 [3]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)

	fmt.Println(reflect.TypeOf(arr3))

	printArray(arr2)
	fmt.Println(arr2)

	// 指针数组
	printArrayByPoint(&arr2)
	fmt.Println(arr2)


}

// 这里只能接收  [3]int的数组
// [10]int  [20]int是不同的类型
// 调用会拷贝数组  数组的传递是值类型传递
func printArray(arr [3]int)  {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, "=", v)
	}
}

// 传指针进行值得修改
func printArrayByPoint(arr *[3]int)  {
	fmt.Println(reflect.TypeOf(arr))
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, "=", v)
	}
}