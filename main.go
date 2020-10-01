package main

import (
	"fmt"
)

/*
func main() {
	//dataStruct.Test()

	//panic2.TestPanic()

	//m_interface.TestInterface()

	// Test goroutine

	//fmt.Println("cpu num : ", runtime.NumCPU())
	//go m_goroutine_channel.Loop()
	//go m_goroutine_channel.Loop()

	//go m_goroutine_channel.Send()
	//go m_goroutine_channel.Receive()
	//time.Sleep(time.Second * 40)


	//m_goroutine_channel.TestWg()

	//m_point.TestPoint()
	//m_point.TestPointArr()

	//m_JSON.TestJson()

	algorithm.Delete()
}
*/

type Student struct {
	Id   int
	Name string
}

func (s Student) Hello(msg string) {
	fmt.Println()
	fmt.Println("hello, ", msg)
}

func Test(maps map[string]interface{}) {
	maps["name"] = "sss"
}
func callback(maps map[string]interface{}, f func(map[string]interface{})) error {
	f(maps)
	return nil
}
func main() {
	//m_interface.TestTypeAssert()
	//u := Student{
	//	Id:   1,
	//	Name: "反射",
	//}
	//u.Hello("1212")

	//v := m_reflect.ValueOf(u)
	//
	//// 获取方法控制权 返回v的名为name的方法的已绑定（到v的持有值的）状态的函数形式的Value封装
	//method := v.MethodByName("Hello")
	//// 拼凑参数
	//args := []m_reflect.Value{m_reflect.ValueOf("反射")}
	//// 调用函数
	//method.Call(args)
}
