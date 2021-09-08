package tree

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) NewTreeNode(val int) *TreeNode {
	tn.Val = val
	tn.Left = nil
	tn.Right = nil
	return tn
}

func InitBinaryTree(root *TreeNode) *TreeNode {
	l := TreeNode{}
	r := TreeNode{}
	root.Left = l.NewTreeNode(9)
	root.Right = r.NewTreeNode(20)
	rl := TreeNode{}
	rr := TreeNode{}
	root.Right.Left = rl.NewTreeNode(15)
	root.Right.Right = rr.NewTreeNode(7)
	return root
}

func BinaryTree() *TreeNode {
	root := &TreeNode{}
	node := InitBinaryTree(root.NewTreeNode(3))
	return node
}

func GetNodeSum(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return GetNodeSum(root.Left) + GetNodeSum(root.Right) + 1
	}
}

func GetDegree(root *TreeNode) (maxDegree int) {
	if root == nil {
		return 0
	}

	maxDegree = 0
	leftMax, rightMax := GetDegree(root.Left), GetDegree(root.Right)
	if leftMax > rightMax {
		maxDegree = leftMax
	} else {
		maxDegree = rightMax
	}
	return maxDegree + 1
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d->", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}

	InOrder(root.Left)
	fmt.Printf("%d->", root.Val)
	InOrder(root.Right)
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d->", root.Val)
}

func IsBlanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ldeepth := GetDegree(root.Left)
	rdeepth := GetDegree(root.Right)

	blanced := false
	if math.Abs(float64(ldeepth-rdeepth)) <= 1 {
		blanced = true
	} else {
		blanced = false
	}
	return blanced && IsBlanced(root.Left) && IsBlanced(root.Right)
}

func LevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	q := []*TreeNode{root}

	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
