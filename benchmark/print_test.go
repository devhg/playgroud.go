package benchmark

import "testing"

// 利用原子实现并发同步
func BenchmarkAtomicMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AtomicMain()
	}
}

func BenchmarkChanMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChanMain()
	}
}

func BenchmarkCtxMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CtxMain()
	}
}
