package main

import (
	"strings"
	"testing"
)

const (
	intput1 = "1cbebd5d-f93a-4b3d-a866-9f4d7e4db3c4,97f36c1b-a867-4ac2-ad44-0d3f1769241b,29ab30a9-8aa2-4d47-ad2f-d9fc42446ecd,9723b38c-639d-4fbe-9772-b3f739797b47"
	intput2 = "a976afec-7d85-4223-a755-2b6c01859f59,5655fa69-90d5-425d-a2d0-e3ac649e2ba8,97f36c1b-a867-4ac2-ad44-0d3f1769241b,29ab30a9-8aa2-4d47-ad2f-d9fc42446ecd"
	intput3 = "4b2e6bcd-efd0-48ee-be3b-192b797c048c,375a5a8a-f844-47ec-9ec1-72d022a665d3,a4b7d781-f118-4a5d-92bf-87bacd9e1b44,3c3d62cf-0cd9-410e-922c-63433ca01e43"
	target  = "97f36c1b-a867-4ac2-ad44-0d3f1769241b"
)

func findWithContains(input, s string) bool {
	return strings.Contains(input, s)
}

func findWithSplit(input, s string) bool {
	parts := strings.Split(input, ",")

	for _, part := range parts {
		if s == part {
			return true
		}
	}

	return false
}

func BenchmarkWithContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findWithContains(intput1, target)
		findWithContains(intput2, target)
		findWithContains(intput3, target)
	}
}

func BenchmarkWithSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findWithSplit(intput1, target)
		findWithSplit(intput2, target)
		findWithSplit(intput3, target)
	}
}
