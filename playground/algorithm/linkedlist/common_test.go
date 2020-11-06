package linkedlist

import (
	"fmt"
	"testing"
)

func TestRet2LinkedList(t *testing.T) {
	l1, l2 := Ret2LinkedList(2, 3)
	l1.String()
	l2.String()

	addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) {
	aa := make([]int, 0)
	fmt.Printf(" aa=%p\n", aa)
	fmt.Printf("&aa=%p\n", &aa)
	pushStack_(&aa)
	fmt.Println(aa)
}

func pushStack_(a *[]int) {
	fmt.Printf("*a=%p\n", *a) // a值的寻址 *(保存的地址) = *(&aa) =  aa
	fmt.Printf(" a=%p\n", a)  // a的值(保存的地址) 实参&aa
	fmt.Printf("&a=%p\n", &a) // 形参指针的地址
	*a = append(*a, 12)
}
