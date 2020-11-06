package m_interface

/*
学习资料

interface 分为iface(非空接口)   eface(空接口)两种
对于eface空接口类型，他会有一个_type字段记录该interface具体的类型，
一旦interface被赋值了一个struct，就算struct是nil，那该_type就是
struct类型。如果_type不是nil，那么该interface也不是nil

只需要记住，interface 当 _type 和 data 同时为nil才是nil就行
*/
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
