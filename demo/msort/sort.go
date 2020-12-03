package msort

import "sort"

// General sort function
func Sort(data sort.Interface) {
	sort.Sort(data)
}

// Test if data is sorted
func IsSorted(data sort.Interface) bool {
	return sort.IsSorted(data)
}

func Reverse(data sort.Interface) sort.Interface {
	return sort.Reverse(data)
}

// Convenience types for common cases which implement msort.Interface
type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Float64Array []float64

func (p Float64Array) Len() int           { return len(p) }
func (p Float64Array) Less(i, j int) bool { return p[i] < p[j] }
func (p Float64Array) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type StringArray []string

func (p StringArray) Len() int           { return len(p) }
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Convenience wrappers for common cases
func SortInts(a []int)         { Sort(IntArray(a)) }
func SortFloat64s(a []float64) { Sort(Float64Array(a)) }
func SortStrings(a []string)   { Sort(StringArray(a)) }

func IntsAreSorted(a []int) bool         { return IsSorted(IntArray(a)) }
func Float64sAreSorted(a []float64) bool { return IsSorted(Float64Array(a)) }
func StringsAreSorted(a []string) bool   { return IsSorted(StringArray(a)) }
