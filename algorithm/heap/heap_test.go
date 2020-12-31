package heap_test

import (
	"container/heap"
	"fmt"
	"testing"
)

//type IntHeap []int
//
//func (h IntHeap) Len() int {
//	return len(h)
//}
//
//func (h IntHeap) Less(i, j int) bool {
//	//return h[i] < h[j] // 小顶堆
//	return h[i] > h[j] // 大顶堆
//}
//
//func (h IntHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *IntHeap) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//
//func (h *IntHeap) Pop() interface{} {
//	n := len(*h)
//	x := (*h)[n-1]
//	*h = (*h)[:n-1]
//	return x
//}

type hp []int

func (h hp) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Len() int {
	return len(h)
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *hp) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

func TestRun(t *testing.T) {
	nums := []int{10, 4, 2, 10}
	h := hp(nums)
	heap.Init(&h)
	heap.Push(&h, 3)
	fmt.Printf("maxnum:%d\n", h[0])
	for h.Len() > 0 {
		fmt.Println(heap.Pop(&h))
	}
}
