package m_common

import (
	"fmt"
	"reflect"
	"runtime"
)

/**
1、返回值类型写在最后面
2、可以返回多个值
3、函数可以作为参数
4、没有默认参数，可选参数
5、有可变长参数
 */

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()  // 反射获取函数指针
	name := runtime.FuncForPC(pointer).Name() //获取函数的全名
	fmt.Printf("Calling %s with args (%d, %d)\n", name, a, b)
	return op(a, b)
}

// 函数作为参数传递
func Test_func_params() {
	result := apply(func(i1 int, i2 int) int {
		return i1 + i2
	}, 1, 2)

	fmt.Println(result)
}

// 可变长参数
func Test_longer_args(nums ...int) (sum int){
	//sum = 0
	for i := range nums {
		sum += nums[i]
	}
	return
}