package stack

import "fmt"

type Stack []interface{}

func (stac *Stack) Push(val interface{}) {
	fmt.Printf("%T\n", stac)
	*stac = append(*stac, val)
}

func (stac *Stack) Pop() interface{} {
	return (*stac)[0]
}
