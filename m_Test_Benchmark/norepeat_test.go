package m_Test_Benchmark

import "testing"

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
		if got := cn_lengthOfLongestSubstring(tt.s); got != tt.ans {
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
	if got := cn_lengthOfLongestSubstring(s); got != ans {
		b.Errorf("cn_lengthOfLongestSubstring() = %v, want %v", got, ans)
	}
}
