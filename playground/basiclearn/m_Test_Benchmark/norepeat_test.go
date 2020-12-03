package m_Test_Benchmark

import (
	"fmt"
	"testing"
)

//测试
// go test .
// go test -coverprofile=c.out
// go tool cover -html=c.out
func Test_cn_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	// 表格测试
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"", 0},
		{"b", 12},
		{"bbbbbbbb", 1},
		{"我是你我是", 3},
	}
	for _, tt := range tests {
		if got := Cn_lengthOfLongestSubstring(tt.s); got != tt.ans {
			t.Errorf("cn_lengthOfLongestSubstring() = %v, want %v", got, tt.ans)
		}
	}
}

//性能基准测试
// 1000000000次	         0.000005 ns
//go test -bench .
func Benchmark_cn_lengthOfLongestSubstring(b *testing.B) {
	s := "按实际代理库时间时收集到的"
	ans := 8
	if got := Cn_lengthOfLongestSubstring(s); got != ans {
		b.Errorf("cn_lengthOfLongestSubstring() = %v, want %v", got, ans)
	}
}

//实用pprof进行性能调优
//go test -bench . -cpuprofile cpu.out
//go tool pprof cpu.out

func ExampleHello() {
	fmt.Println(112)
	// Output:
	// 112
}

func Fib(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
