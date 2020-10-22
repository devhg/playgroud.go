package tree

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func (tn *TreeNode) NewTreeNode(val int) *TreeNode {
	tn.Value = val
	tn.LeftNode = nil
	tn.RightNode = nil
	return tn
}

func InitBinaryTree(root *TreeNode) *TreeNode {
	l := TreeNode{}
	r := TreeNode{}
	root.LeftNode = l.NewTreeNode(2)
	root.RightNode = r.NewTreeNode(3)
	l2 := TreeNode{}
	ll2 := TreeNode{}
	root.LeftNode.LeftNode = l2.NewTreeNode(4)
	root.LeftNode.RightNode = ll2.NewTreeNode(5)
	return root
}

func BinaryTree() *TreeNode {
	root := TreeNode{}
	node := InitBinaryTree(root.NewTreeNode(1))
	return node
}

func GetNodeSum(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return GetNodeSum(root.LeftNode) + GetNodeSum(root.RightNode) + 1
	}
}

func GetDegree(root *TreeNode) (maxDegree int) {
	if root == nil {
		return 0
	}

	maxDegree = 0
	leftMax, rightMax := GetDegree(root.LeftNode), GetDegree(root.RightNode)
	if leftMax > rightMax {
		maxDegree = leftMax
	} else {
		maxDegree = leftMax
	}
	return maxDegree + 1
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d->", root.Value)
	PreOrder(root.LeftNode)
	PreOrder(root.RightNode)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}

	InOrder(root.LeftNode)
	fmt.Printf("%d->", root.Value)
	InOrder(root.RightNode)
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	PostOrder(root.LeftNode)
	PostOrder(root.RightNode)
	fmt.Printf("%d->", root.Value)
}

func IsBlanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ldeepth := GetDegree(root.LeftNode)
	rdeepth := GetDegree(root.RightNode)

	blanced := false
	if math.Abs(float64(ldeepth-rdeepth)) <= 1 {
		blanced = true
	} else {
		blanced = false
	}
	return blanced && IsBlanced(root.LeftNode) && IsBlanced(root.RightNode)
}
