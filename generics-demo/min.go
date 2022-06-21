package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func logMin[T constraints.Ordered](x, y T) {
	fmt.Printf("min(%v, %v) = %v\n", x, y, min(x, y))
}

type MyInt int

func TestMin() {
	var i1 MyInt = 7
	var i2 MyInt = 8

	logMin(5, 7)
	logMin(5.4, 7.6)
	logMin(i1, i2)
	logMin("a", "b")
	// Output:
	// min(5, 7) = 5
	// min(5.4, 7.6) = 5.4
	// min(7, 8) = 7
	// min(a, b) = a
}
