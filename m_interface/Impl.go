package m_interface

import (
	"fmt"
)

/*
关于实现接口：
1. 一个struct可以实现多个接口
2. 接口间可以相互组合
3. 接受者类型必须保持一致，要么全是指针接收者  要么全是值接收者
*/

type Cat struct {
	ID   int
	Name string
	Age  int
}

type Dog struct {
	ID   int
	Name string
	Age  int
}

type mPig struct {
}

func (m mPig) String() string {
	// 实现stringer 相当于java的toString 方法
	// fmt.Println(bp)会打印返回值
	return "{mPig:...}"
}

func (m mPig) Run() string {
	panic("implement me")
}

func (m mPig) Eat() string {
	panic("implement me")
}

func (m mPig) Sleep() string {
	panic("implement me")
}

func (m mPig) Drink() string {
	panic("implement me")
}

func (c *Cat) Eat() string {
	fmt.Println("cat eating")
	return ""
}

func (c *Cat) Run() string {
	fmt.Println("cat Running")
	return ""
}

func (d *Dog) Eat() string {
	fmt.Println("dog eating")
	return ""
}

func (d *Dog) Run() string {
	fmt.Println("dog Running")
	return ""
}

func AnimalDo(a Animal) {
	a.Eat()
	a.Run()
}

func PigDo(bp BigPig) {
	bp.Eat()
	bp.Run()
	bp.Drink()
	bp.Sleep()
}
func TestInterface() {
	cat := Cat{
		ID:   1,
		Name: "11",
		Age:  0,
	}

	dog := Dog{
		ID:   2,
		Name: "22",
		Age:  0,
	}

	var bp BigPig
	bp = mPig{}
	fmt.Printf("%T, %v\n", bp, bp)

	var p Pig
	p = bp
	fmt.Printf("%T, %v\n", p, p)

	switch v := p.(type) {
	case *mPig:
		fmt.Println("*mPig", v)
	case mPig:
		fmt.Println("mPig")
	default:
		break
	}

	fmt.Println(&cat)
	AnimalDo(&cat)
	AnimalDo(&dog)

	fmt.Println(bp)
}
