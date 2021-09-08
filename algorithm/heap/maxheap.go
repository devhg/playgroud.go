package heap

// import "fmt"

// type maxHeap struct {
// 	k    int
// 	heap []int
// }

// func NewMaxHeap(k int, nums []int) *maxHeap {
// 	heap := &maxHeap{k: k, heap: []int{}}
// 	for _, n := range nums {
// 		heap.Push(n)
// 	}
// 	return heap
// }

// func (this *maxHeap) Push(num int) {
// 	if len(this.heap) < this.k {
// 		this.heap = append(this.heap, num)
// 		this.up(len(this.heap) - 1) // 执行上浮，上浮到合适的位置
// 	} else if num > this.heap[len(this.heap)-1] {
// 		this.heap[len(this.heap)-1] = num
// 		this.up(len(this.heap) - 1)
// 	}
// }

// func (this *maxHeap) Pop() (ret int) {
// 	ret = this.heap[0]
// 	this.heap = this.heap[1:]
// 	this.down(0)
// 	return
// }

// func (this *maxHeap) up(i int) {
// 	for i > 0 { // 上浮到索引0就停止上浮，0是堆顶位置
// 		parent := (i - 1) >> 1 // 找到父节点在heap数组中的位置
// 		if this.heap[parent] < this.heap[i] {
// 			this.heap[parent], this.heap[i] = this.heap[i], this.heap[parent] // 交换
// 			i = parent                                                        // 更新 i
// 		} else {
// 			break
// 		}
// 	}
// }

// func (this *maxHeap) down(i int) { // 下沉到合适的位置
// 	for 2*i+1 < len(this.heap) { // 左子节点的索引如果已经越界，终止下沉
// 		child := 2*i + 1
// 		if child+1 < len(this.heap) && this.heap[child+1] > this.heap[child] {
// 			child++ // 左右孩子中 取较大的 去比较
// 		}
// 		if this.heap[i] < this.heap[child] { // 如果插入的数字比子节点都小
// 			this.heap[child], this.heap[i] = this.heap[i], this.heap[child] // 交换
// 			i = child                                                       // 更新 i
// 		} else {
// 			break
// 		}
// 	}
// }

// func main() {
// 	heap := NewMaxHeap(3, []int{3, 4, 2})
// 	heap.Push(11)
// 	for len(heap.heap) > 0 {
// 		fmt.Println(heap.Pop())
// 	}
// 	//fmt.Println(heap.heap)
// 	//heap.Remove()
// }
