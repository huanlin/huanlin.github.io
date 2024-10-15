package main

import (
	"testing"
)

func BenchmarkShuffle1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shuffle1("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
}

func BenchmarkShuffle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shuffle2("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
}
