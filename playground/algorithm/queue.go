package algorithm

type Queue []interface{}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	// xxx.(int) 强制类型转换
	return head.(int)
}
