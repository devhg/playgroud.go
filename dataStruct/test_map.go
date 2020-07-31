package dataStruct

import (
	"fmt"
	"sort"
)

//map使用哈希表，必须可以比较相等
//除了slice map function的内建类型都可以作为key
//struct类型不包含上述字段，也可以作为可以
func TestMap() {
	fmt.Println([]byte("abdbcd"))
	//defineMap()
	//traverMap()
	//getValue()
	//deleteMap()
	//traverMapSorted()
}

// 定义map
func defineMap() {
	var m1 = map[string]interface{}{
		"name": "222",
		"age":  12,
	}

	m2 := make(map[string]int) // m2==empty map
	var m3 map[int]string      // m3 == nil

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3 == nil)
}

// 遍历Map
func traverMap() {
	fmt.Println("Travering map")
	var m1 = map[string]interface{}{
		"name": "222",
		"age":  12,
	}

	for k := range m1 {
		println(k)
	}

	for _, v := range m1 {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k, " ", v)
	}
}

// 顺序遍历map
// map遍历是无序的
//如果需要顺序，需要通过切片对key进行手动排序
func traverMapSorted() {
	fmt.Println("Travering map sorted")
	var m1 = map[string]interface{}{
		"name":   "222",
		"age":    12,
		"course": "golang",
	}

	var keys []string
	for key := range m1 {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m1[k])
	}
}

// 获取元素
func getValue() {
	fmt.Println("Geting values")
	var m1 = map[string]interface{}{
		"name": "222",
		"age":  12,
	}

	name := m1["name"]
	fmt.Println(name)
	course := m1["course"] // 若没有该key，map返回的是map的v类型的零值 interface->nil  int->0  string->""
	fmt.Println(course)

	name, ok := m1["name"]
	fmt.Println(name, ok)
	if course, ok := m1["course"]; ok {
		fmt.Println(course, ok)
	} else {
		fmt.Println("key not found")
	}
}

// 删除元素
func deleteMap() {
	fmt.Println("Deleting values")
	var m1 = map[string]interface{}{
		"name": "222",
		"age":  12,
	}

	name, ok := m1["name"]
	fmt.Println(name, ok)

	delete(m1, "name")

	name, ok = m1["name"]
	fmt.Println(name, ok)
}
