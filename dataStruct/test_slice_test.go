package dataStruct

import (
	"fmt"
	"testing"
)

func TestTestSlice(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:8]

	//s2 = append(s2, 100)
	//s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	//fmt.Println(s2)
	//fmt.Println(slice)
}

func appe(data []int) {
	data = append(data, 1222)
}

// slice := []int{0,1,2,3,4,5,6,7,8,9}
// s1 := slice[2:5]  len=5-2=3  cap=cap(slice)-2=8
// s2 := s1[2:x:y]  len=x-2  cap=y-2<=8  y不能超过被切片的容量的索引

func TestTestSlice2(t *testing.T) {
	slice := []int{0, 1, 2}
	fmt.Println(len(slice), cap(slice))
	appe(slice)
	fmt.Println(slice)
}
