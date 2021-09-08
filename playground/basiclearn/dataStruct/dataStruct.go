package datastruct

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) Hello() {

}

func Test() {

	// make返回引用类型
	// new 返回指针类型

	mMake := make(map[int]int)
	mNew := new(map[int]int)

	fmt.Println("mMake:", reflect.TypeOf(mMake)) //mMake: map[int]int
	fmt.Println("mNew:", reflect.TypeOf(mNew))   //mNew: *map[int]int

	// append delete  copy

	strings := make([]string, 2)
	strings[0] = "id-1"
	strings[1] = "id-2"
	strings = append(strings, "id-3")

	fmt.Println("len:", len(strings)) //len: 3
	fmt.Println("cap:", cap(strings)) //len: 4

	mMake[1] = 111
	mMake[2] = 222
	mMake[3] = 333

	fmt.Println(mMake) //map[1:111 2:222 3:333]
	delete(mMake, 2)
	fmt.Println(mMake) //map[1:111 3:333]
}
