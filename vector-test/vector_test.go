package main

import "testing"

var (
	testData = []Vector3D{
		{1000.123, 2000.234, 3000.345},
		{2000.123, 3000.234, 7000.345},
		{3000.123, 4000.234, 8000.345},
		{4000.123, 5000.234, 9000.345},
		{5000.123, 6000.234, 1000.345},
		{-100.123, -200.234, 3000.345},
		{-200.123, -300.234, 7000.345},
		{-300.123, -400.234, 8000.345},
		{-400.123, -500.234, 9000.345},
		{-500.123, -600.234, 1000.345},
	}
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vector := testData[0]
		for _, v := range testData {
			vector.Add(&v)
		}
	}
}

func BenchmarkAddByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vector := testData[0]
		for _, v := range testData {
			vector.AddByValue(v)
		}
	}
}

func BenchmarkAddRetunCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vector := testData[0]
		for _, v := range testData {
			vector = AddReturnCopy(vector, v)
		}
	}
}
