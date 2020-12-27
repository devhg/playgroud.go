package linkedlist

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 传入l1节点数量，l2节点数量
func Ret2LinkedList(a, b int) (l1, l2 *ListNode) {
	l1, l2 = &ListNode{}, &ListNode{}
	cur := l1
	for i := 0; i < a; i++ {
		node := &ListNode{
			Val:  rand.Intn(10),
			Next: nil,
		}
		cur.Next = node
		cur = node
	}

	cur = l2
	for i := 0; i < b; i++ {
		node := &ListNode{
			Val:  rand.Intn(10),
			Next: nil,
		}
		cur.Next = node
		cur = node
	}

	return l1.Next, l2.Next
}

// 传入l节点数量
func Ret1LinkedList(a int) (l1 *ListNode) {
	l1 = &ListNode{}
	cur := l1
	for i := 0; i < a; i++ {
		node := &ListNode{
			Val:  rand.Intn(10),
			Next: nil,
		}
		cur.Next = node
		cur = node
	}

	return l1.Next
}

// 由slice构建链表
func NewLinkedList(a []int) *ListNode {
	l := &ListNode{}
	cur := l
	for _, i := range a {
		node := &ListNode{
			Val:  i,
			Next: nil,
		}
		cur.Next = node
		cur = node
	}
	return l.Next
}

func (l *ListNode) String() {
	data := make([]int, 0)
	for l != nil {
		data = append(data, l.Val)
		l = l.Next
	}
	fmt.Println(data)
}
