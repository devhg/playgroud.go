package skiplist

import (
	"fmt"
	"math/rand"
)

const (
	MaxLevel    = 16  // 最大层数限制
	LevelFactor = 0.5 // 定义生成层级的因子
)

// 几种错误码
const (
	Ok = iota + 1
	Duplicated
	NotExist
	NotInit
)

// 定义一个描述数据的接口
// 用来描述一个实际存储的对象。
type Interface interface {
	Less(p Interface) bool
	Equal(p Interface) bool
}

type fakeNode struct {
	// 哨兵节点，伪节点
}

func (f fakeNode) Less(p Interface) bool {
	return false
}

func (f fakeNode) Equal(p Interface) bool {
	return false
}

// 链表节点
type node struct {
	data     Interface // 实际的数据
	forwards []*node   // 索引存储的位置
	level    int       // 节点的层级
}

// 链表
type SkipList struct {
	head  *node  // 链表头节点，存储fakeNode
	len   uint32 // 链表长度
	level int    // 当前跳表最高层级
}

// 初始化一个节点
func newNode(p Interface, l int) *node {
	return &node{data: p, forwards: make([]*node, l), level: l}
}

// 初始化一个跳表
func NewSkipList() *SkipList {
	return &SkipList{head: newNode(&fakeNode{}, MaxLevel), len: 0, level: 1}
}

// 这里随机产生一个 层级
// 在 LEVEL_FACTOR 是 0.5 的情况下
// 1 级的概率是 50%
// 2 级的概率是 25%
// 3 级的概率是 12.5%, 以此类推
func (sl *SkipList) randomLevel() int {
	l := 1

	for rand.Float64() < LevelFactor && l < MaxLevel {
		l++
	}

	// 如果层级比当前层级高2级或以上，按照高一级处理，避免浪费
	if sl.level+1 < l {
		return sl.level + 1
	}
	return l
}

func (sl *SkipList) Add(p Interface) int {
	// 如果 head 为空，说明没有初始化
	if sl.head == nil {
		return NotInit
	}

	cur := sl.head                // 指向哨兵节点，从最高层开始查找插入为止
	roadNode := [MaxLevel]*node{} // 跳跃路上的节点

	for i := MaxLevel - 1; i >= 0; i-- {
		// 当前层级向后查找，找到当前节点下一个是nil 或 这下一个节点值大于p 为止
		for ; cur.forwards[i] != nil; cur = cur.forwards[i] {
			if cur.forwards[i].data.Equal(p) {
				return Duplicated
			}
			// p小于下一个节点的data
			if !cur.forwards[i].data.Less(p) {
				roadNode[i] = cur // 记录下当前节点
				break
			}
		}

		// 如果这个层级遍历结束，还是没有找到对应位置
		// 那么就将最后的元素作为当前层级路径
		if cur.forwards[i] == nil {
			roadNode[i] = cur
		}
	}

	sl.len++ // 找到插入位置，跳表长度加一

	l := sl.randomLevel() // 生成新节点层级
	n := newNode(p, l)    // 初始化节点

	// 从最底层开始添加节点和索引，依据来时的路径
	for i := 0; i < n.level; i++ {
		next := roadNode[i].forwards[i]
		n.forwards[i] = next
		roadNode[i].forwards[i] = n
	}

	// 更新跳表的索引最高层级
	if n.level > sl.level {
		sl.level = n.level
	}
	return Ok
}

func (sl *SkipList) Delete(p Interface) int {
	cur := sl.head
	roadNode := [MaxLevel]*node{}

	for i := sl.level - 1; i >= 0; i-- {
		roadNode[i] = cur
		for ; cur.forwards[i] != nil; cur = cur.forwards[i] {
			if !cur.forwards[i].data.Less(p) {
				roadNode[i] = cur
				break
			}
		}
	}
	cur = roadNode[0].forwards[0] // 要删的节点
	if cur == nil {
		return NotExist
	}

	// 从cur节点所在的最高层级依次向下删除(必须)
	for i := cur.level - 1; i >= 0; i-- {
		// 如果当前节点是某一个层级的最后一个元素
		// 那么降低跳表的层级
		if roadNode[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}
		// 删除元素
		if roadNode[i].forwards[i] != nil {
			roadNode[i].forwards[i] = roadNode[i].forwards[i].forwards[i]
		}
	}

	sl.len--
	return Ok
}

func (sl *SkipList) Search(p Interface) *node {
	cur := sl.head

	for i := sl.level; i >= 0; i-- {
		for ; cur.forwards[i] != nil; cur = cur.forwards[i] {
			if cur.forwards[i].data.Equal(p) {
				return cur.forwards[i]
			}
			if !cur.forwards[i].data.Less(p) {
				break
			}
		}
	}
	return nil
}

func (sl *SkipList) Length() uint32 {
	return sl.len
}

func (sl *SkipList) Level() int {
	return sl.level
}

func (sl *SkipList) Print() {
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		fmt.Printf("[level = %d] ", i)
		for nil != cur {
			fmt.Printf("%+v   ", cur.data)
			cur = cur.forwards[i]
		}
		fmt.Println()
		cur = sl.head
	}
}
