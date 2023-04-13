package main

import "testing"

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort([]int{45, 23, 17, 38, 31})
	}
}
