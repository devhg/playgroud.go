package m_interface

import "fmt"

type Animal interface {
	Run() string
	Eat() string
}

type Pig interface {
	Sleep() string
}

type BigPig interface {
	Animal
	Pig
	Drink() string
}

type name struct {
	Age string
}

//下面方法只适合于interface
//xx.(type)只能用于switch中
//xx.(int)强制转成int类型
func TestTypeAssert() {
	var q interface{}
	q = &name{Age: "sss"}

	switch v := q.(type) {
	case name:
		fmt.Println("name", v)
	case *name:
		fmt.Println("*name", v)
	}

	if s, ok := q.(int); ok {
		fmt.Println(s)
	} else {
		fmt.Println("errr")
	}

}
