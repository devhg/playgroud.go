package linkedlist

import (
	"fmt"
	"testing"
)

func TestRet2LinkedList(t *testing.T) {
	//list2 := NewLinkedList([]int{1, 4, 3, 2, 5, 2})
	//list2.String()
	////*list2 = *list2.Next
	//list2 = partition(list2, 3)
	//list2.String()
	fmt.Println((*ListNode)(nil) == (*ListNode)(nil))
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		// 如果第一次遍历到链表尾部，就指向另一个链表的头部，
		//继续遍历，这样会抵消长度差。如果没有相交，因为遍历长度相等，最后会是 nil ==  nil
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
