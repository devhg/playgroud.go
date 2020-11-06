package m_reflect

import (
	"testing"
)

func TestReflect_(t *testing.T) {
	//TestReflect1()
	//modifyByReflect() // 测试修改
	//TestReflectField()//测试字段反射

	//TestCallMethod() // 反射值接收器

	//TestCallMethodPtr() // 反射指针接收器

	//TestBack() // 测试反射对象还原

	TestUnknown(User{
		Name: "1",
		Age:  0,
		From: "2",
	}) // 测试未知类型
}
