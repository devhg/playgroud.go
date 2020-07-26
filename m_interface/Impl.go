package m_interface

import "fmt"

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

func AnimalDo(a Animal)  {
	a.Eat()
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

	fmt.Println(&cat)
	AnimalDo(&cat)
	AnimalDo(&dog)
}
