package main

import (
	"testing"
)

var testData = []Point{
	{1000.124, 2000.234, 3000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{2000.124, 3000.234, 7000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{3000.124, 4000.234, 8000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{4000.124, 5000.234, 9000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{5000.124, 6000.234, 1000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{-100.124, -200.234, 3000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{-200.124, -300.234, 7000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{-300.124, -400.234, 8000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{-400.124, -500.234, 9000.345, "a", "b", "c", 1, 2, 3, false, true, false},
	{-500.124, -600.234, 1000.345, "a", "b", "c", 1, 2, 3, false, true, false},
}

func Benchmark_Add(b *testing.B) {
	var point Point

	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			point.Add(&v)
		}
	}

	if point.x < 0 || point.y < 0 || point.z < 0 {
		b.Fatalf("unexpected result: %v", point)
	}
}

func Benchmark_AddByValue(b *testing.B) {
	var point Point

	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			point.AddByValue(v)
		}
	}

	if point.x < 0 || point.y < 0 || point.z < 0 {
		b.Fatalf("unexpected result: %v", point)
	}
}

func Benchmark_Add_InPlace(b *testing.B) {
	var point Point

	for i := 0; i < b.N; i++ {
		point.Add(&Point{
			x:    0.245,
			y:    45.345345,
			z:    6456.2342,
			str1: "foo",
			str2: "bar",
			str3: "baz",
			i1:   0,
			i2:   1,
			i3:   2,
			b1:   false,
			b2:   true,
			b3:   false,
		})
	}

	if point.x < 0 || point.y < 0 || point.z < 0 {
		b.Fatalf("unexpected result: %v", point)
	}
}

func Benchmark_AddByValue_InPlace(b *testing.B) {
	var point Point

	for i := 0; i < b.N; i++ {
		point.AddByValue(Point{
			x:    0.245,
			y:    45.345345,
			z:    6456.2342,
			str1: "foo",
			str2: "bar",
			str3: "baz",
			i1:   0,
			i2:   1,
			i3:   2,
			b1:   false,
			b2:   true,
			b3:   false,
		})
	}

	if point.x < 0 || point.y < 0 || point.z < 0 {
		b.Fatalf("unexpected result: %v", point)
	}
}

func Benchmark_AddReturnCopy(b *testing.B) {
	var point Point

	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			point = AddReturnCopy(point, v)
		}
	}

	if point.x < 0 || point.y < 0 || point.z < 0 {
		b.Fatalf("unexpected result: %v", point)
	}
}
