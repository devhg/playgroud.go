package strcontact_test

import (
	"github.com/devhg/LearnGo/benchmark/strcontact"
	"testing"
)

func BenchmarkTestAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByAdd()
	}
}

func BenchmarkTestFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByfmSpf()
	}
}

func BenchmarkTestBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByBuffer()
	}
}

func BenchmarkTestAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcontact.StrContactByAppend()
	}
}
