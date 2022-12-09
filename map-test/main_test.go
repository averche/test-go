package main

import (
	"testing"
)

func BenchmarkNumberToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberToRoman(1066)
	}
}

func BenchmarkNumberToRomanGlobalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberToRomanGlobalMap(1066)
	}
}
