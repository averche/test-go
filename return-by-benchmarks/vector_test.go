package main

import (
	"testing"
)

var (
	testData = []Vector3D{
		{1000.124, 2000.234, 3000.345},
		{2000.124, 3000.234, 7000.345},
		{3000.124, 4000.234, 8000.345},
		{4000.124, 5000.234, 9000.345},
		{5000.124, 6000.234, 1000.345},
		{-100.124, -200.234, 3000.345},
		{-200.124, -300.234, 7000.345},
		{-300.124, -400.234, 8000.345},
		{-400.124, -500.234, 9000.345},
		{-500.124, -600.234, 1000.345},
	}
)

func BenchmarkAdd(b *testing.B) {
	var vector Vector3D

	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			vector.Add(&v)
		}
	}
}

func BenchmarkAddByValue(b *testing.B) {
	var vector Vector3D

	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			vector.AddByValue(v)
		}
	}
}

func BenchmarkAddReturnCopy(b *testing.B) {
	var vector Vector3D

	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			vector = AddReturnCopy(vector, v)
		}
	}
}
