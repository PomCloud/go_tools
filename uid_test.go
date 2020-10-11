package go_tools

import "testing"

func BenchmarkRandString(b *testing.B) {
	n := 0x08
	for i := 0; i < b.N; i++ {
		RandString(n)
	}
}
