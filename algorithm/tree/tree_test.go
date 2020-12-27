package tree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	PreOrder(tree)
}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}

	for _, num := range nums {
		root = insert(num, root)
	}
	return root
}
func insert(val int, root *TreeNode) *TreeNode {
	if val < root.Val {
		insertRight(val, root)
		return root
	}
	return reRoot(val, root)
}

func insertRight(val int, root *TreeNode) {
	cur := root
	for {
		if cur.Right != nil {
			if cur.Right.Val < val {
				cur.Right = reRoot(val, cur.Right)
				return
			}
			cur = cur.Right
		} else {
			cur.Right = &TreeNode{Val: val}
			return
		}
	}
}

func reRoot(val int, root *TreeNode) (node *TreeNode) {
	node = &TreeNode{Val: val, Left: nil, Right: nil}
	node.Left = root
	return
}
