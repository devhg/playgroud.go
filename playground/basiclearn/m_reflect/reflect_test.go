package m_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Users struct {
	Name string
	Age  int
}

//func (u *Users) M1() error {
//	fmt.Println("M1")
//	return nil
//}

func (u *Users) M2(num int) {
	fmt.Println("M2", num)
}

func reflects(u interface{}) {
	indirect := reflect.Indirect(reflect.ValueOf(u))
	t := indirect.Type()
	fmt.Println(t.NumMethod())
	//oType := reflect.TypeOf(u)
	//println(oType.Name())
	//fmt.Println(oType.Kind())
	//
	//fmt.Println(oType.NumField())
	//fmt.Println(oType.NumMethod())

	//field := oType.Field(0)
	//fmt.Println(field)
	//
	//method := oType.Method(0)
	//println(method.Name)
	//methodType := method.Type
	//
	//in := methodType.NumIn()
	//t2 := methodType.In(0)
	//fmt.Println(in)
	//fmt.Println(t2)
}

func callMethod(u interface{}) {
	//rv := reflect.ValueOf(u)
	//t := reflect.TypeOf(u)
	//fmt.Println()
	//fmt.Println(rv.NumMethod())
	//
	//param := []reflect.Value{}
	//for i := 0; i < rv.NumMethod(); i++ {
	//	method := rv.Method(i)
	//	t := method.Type()
	//
	//	for j := 0; j < t.NumIn(); j++ {
	//		fmt.Println("method in type:", t.In(j))
	//		if t.In(j).Kind() == reflect.Int {
	//			param = append(param, reflect.ValueOf(19))
	//		}
	//		fmt.Println(t.In(j).Name())
	//	}
	//
	//	if call := method.Call(param); len(call) >0 {
	//		fmt.Println(call)
	//	}
	//}

}

func TestReflect_(t *testing.T) {
	//CallMethod("M1", &Users{})
	//var i interface{} = &Users{}
	callMethod(Users{})

}
