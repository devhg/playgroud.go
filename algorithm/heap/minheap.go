package heap

type MinHeap struct {
	k    int
	heap []int
}

func NewMinHeap(k int, nums []int) *MinHeap {
	heap := &MinHeap{k: k, heap: []int{}}
	for _, n := range nums {
		heap.Push(n)
	}
	return heap
}

// Push .
func (h *MinHeap) Push(num int) {
	if len(h.heap) < h.k {
		h.heap = append(h.heap, num)
		h.up(len(h.heap) - 1) // 执行上浮，上浮到合适的位置
	} else if num < h.heap[len(h.heap)-1] { // 如果num比堆顶数字还要大
		h.heap[len(h.heap)-1] = num
		h.up(len(h.heap) - 1)
	}
}

func (h *MinHeap) Pop() (ret int) {
	ret = h.heap[0]
	h.heap = h.heap[1:]
	h.down(0)
	return
}

func (h *MinHeap) up(i int) {
	for i > 0 { // 上浮到索引0就停止上浮，0是堆顶位置
		parent := (i - 1) >> 1          // 找到父节点在heap数组中的位置
		if h.heap[parent] > h.heap[i] { // 如果父节点的数字比插入的数字大
			h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent] // 交换
			i = parent                                            // 更新 i
		} else {
			break
		}
	}
}

func (h *MinHeap) down(i int) { // 下沉到合适的位置
	for 2*i+1 < len(h.heap) { // 左子节点的索引如果已经越界，终止下沉
		// 左子节点在heap数组中的位置
		child := 2*i + 1
		if child+1 < len(h.heap) && h.heap[child+1] < h.heap[child] {
			child++ // 左右孩子中 取较小的 去比较
		}
		if h.heap[i] > h.heap[child] { // 如果插入的数字比子节点都大
			h.heap[child], h.heap[i] = h.heap[i], h.heap[child] // 交换
			i = child                                           // 更新 i
		} else {
			break
		}
	}
}

// func main() {
// 	heap := NewMinHeap(3, []int{3, 4, 2})
// 	heap.Push(1)
// 	for len(heap.heap) > 0 {
// 		fmt.Println(heap.Pop())
// 	}
// 	//fmt.Println(heap.heap)
// 	//heap.Remove()
// }
