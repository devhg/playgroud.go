package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := BinaryTree()

	deepth := IsBlanced(tree)
	fmt.Println(deepth)
	//PreOrder(tree)
}
func TestBinaryTree_(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{5}
	b = append(b, a...)
	fmt.Println(b)
}
