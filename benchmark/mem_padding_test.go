package benchmark

import (
	"sync"
	"testing"
)

// 内存填充 解决多核CPU Cache一致性问题 测试

// 这里M需要足够大，否则会存在goroutine 1已经执行完成，而goroutine 未启动的情况
const M = 1000000

type SimpleStruct struct {
	n int
}

// 使用内存填充
type PaddedStruct struct {
	n int
	_ CacheLinePad
}

type CacheLinePad struct {
	_ [CacheLinePadSize]byte
}

const CacheLinePadSize = 64

func BenchmarkStructureFalseSharing(b *testing.B) {
	structA := SimpleStruct{}
	structB := SimpleStruct{}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < M; j++ {
				structA.n += 1
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < M; j++ {
				structB.n += 1
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkStructurePadding(b *testing.B) {
	structA := PaddedStruct{}
	structB := SimpleStruct{}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < M; j++ {
				structA.n += 1
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < M; j++ {
				structB.n += 1
			}
			wg.Done()
		}()
		wg.Wait()
	}
}
