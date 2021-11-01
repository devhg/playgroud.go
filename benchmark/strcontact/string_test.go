package strcontact_test

import (
	"testing"

	"github.com/devhg/playgroud.go/benchmark/strcontact"
)

// 1000个字符串连接
// goos: darwin
// goarch: amd64
// pkg: github.com/devhg/LearnGo/benchmark/strcontact
// Benchmark名字-CPU          循环次数          平均每次执行时间         平均每次的内存消耗     平均每次的分配次数
// BenchmarkAdd-8              5529            236885 ns/op         1063894 B/op        999 allocs/op
// BenchmarkFormat-8           3289            367180 ns/op         1096736 B/op       3000 allocs/op
// BenchmarkBuffer-8         126193              8971 ns/op            6576 B/op          7 allocs/op
// BenchmarkBuilder-8        161905              6707 ns/op            7416 B/op         11 allocs/op
// BenchmarkAppend-8          88270             13306 ns/op            9464 B/op         12 allocs/op
// PASS
// ok      github.com/devhg/LearnGo/benchmark/strcontact   7.825s
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByAdd()
	}
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByfmSpf()
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByBuffer()
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByBuilder()
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByAppend()
	}
}
