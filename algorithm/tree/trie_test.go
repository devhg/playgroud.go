package tree

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	tree := createTree([]string{"how", "hi", "her", "hello", "so", "see"})
	flag := tree.findWord("her")
	fmt.Println(flag)
	flag = tree.findWord("x")
	fmt.Println(flag)
}
