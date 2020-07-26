package m_reflect

import (
	"fmt"
	"reflect"
)

type order struct {
	id   int
	name string
}
type employee struct {
	name   string
	age    int
	salary float32
}

func insert(t interface{}) {
	if reflect.ValueOf(t).Kind() == reflect.Struct {
		table := reflect.TypeOf(t).Name()
		sql := fmt.Sprintf("insert into %s values(", table)
		v := reflect.ValueOf(t)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					sql = fmt.Sprintf("%s%d", sql, v.Field(i).Int())
				} else {
					sql = fmt.Sprintf("%s, %d", sql, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					sql = fmt.Sprintf("%s%s", sql, v.Field(i).String())
				} else {
					sql = fmt.Sprintf("%s, %s", sql, v.Field(i).String())
				}
			case reflect.Float32:
				if i == 0 {
					sql = fmt.Sprintf("%s%f", sql, v.Field(i).Float())
				} else {
					sql = fmt.Sprintf("%s, %f", sql, v.Field(i).Float())
				}
			default:
				fmt.Println("unsupported type")
			}
		}
		sql = fmt.Sprintf("%s)", sql)
		fmt.Println(sql)
		return
	}
	fmt.Println("unsupported type")
}

type user struct {
	Id   int
	Name string
}

func (u *user) Hello() {
	fmt.Println("hello")
}

func TestReflect(t interface{}) {

	objT := reflect.TypeOf(t)
	objV := reflect.ValueOf(t)

	//获取去这个类型的名称
	fmt.Println("这个类型的名称是:", objT.Name())
	//
	for i := 0; i < objT.NumField(); i++ {
		//从0开始获取结构体所包含的key
		key := objT.Field(i)
		//从0开始通过interface方法来获取key所对应的值
		value := objV.Field(i).Interface()

		fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)
	}

	for i := 0; i < objT.NumMethod(); i++ {
		m := objT.Method(i)
		fmt.Printf("第%d个方法是：%s:%v\n", i+1, m.Name, m.Type)
	}

}

func TestReflect1()  {
	//u := &user{1, "2"}
	//elem := m_reflect.TypeOf(u).Elem()

	//u := user{1, "2"}
	//m_reflect.TypeOf(u)

	//两种等效
}

//func main() {
//	o := order{
//		id:   1,
//		name: "zgh",
//	}
//	//emp := employee{
//	//	name:   "zgh",
//	//	age:    20,
//	//	salary: 1.2,
//	//}
//	insert(o)
//	//insert(emp)
//
//}
