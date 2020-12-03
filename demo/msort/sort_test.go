package msort

import (
	"fmt"
	"testing"
)

// sorting of slice of integers
func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data) //conversion to type IntArray
	Sort(a)
	if !IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
	Sort(Reverse(a))
	fmt.Printf("The reversed array is: %v\n", a)
}

// sorting of slice of strings
func strings() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	a := StringArray(data)
	Sort(a)
	if !IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
	Sort(Reverse(a))
	fmt.Printf("The reversed array is: %v\n", a)
}

// a type which describes a day of the week
type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p dayArray) Len() int           { return len(p.data) }
func (p dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

// sorting of custom type day
func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	Sort(a)
	if !IsSorted(a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}

func TestSort(t *testing.T) {
	ints()
	strings()
	days()
}
