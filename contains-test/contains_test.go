package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func BenchmarkContainsArray(b *testing.B) {
	arr := [...]int{
		123,
		456,
		789,
		234,
		345,
		456,
		567,
		678,
		789,
	}

	for i := 0; i < b.N; i++ {
		if slices.Contains(arr[:], 456) && slices.Contains(arr[:], 789) {
			slices.Contains(arr[:], 123)
		}
	}
}

func BenchmarkContainsSlice(b *testing.B) {
	s := []int{
		123,
		456,
		789,
		234,
		345,
		456,
		567,
		678,
		789,
	}

	for i := 0; i < b.N; i++ {
		if slices.Contains(s, 456) && slices.Contains(s, 789) {
			slices.Contains(s, 123)
		}
	}
}
