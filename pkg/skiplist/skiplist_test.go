package skiplist

import (
	"testing"
)

type Int int

func (i Int) Less(p Interface) bool {
	return i < p.(Int)
}

func (i Int) Equal(p Interface) bool {
	return i == p.(Int)
}

func TestNewSkipList(t *testing.T) {
	skipList := NewSkipList()
	for i := 1; i <= 18; i++ {
		skipList.Add(Int(i))
	}

	skipList.Print()
}
