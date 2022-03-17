package golang_benchmarks

import (
	"testing"
)

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withSwitch("1")
	}
}

func BenchmarkSwitchMiss(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withSwitch("20")
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withMap("1")
	}
}

func BenchmarkMapMiss(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withMap("20")
	}
}
