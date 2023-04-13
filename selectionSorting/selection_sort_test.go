package main

import "testing"

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectionSort([]int{4, 7, 2, 5, 1, 6, 3})
	}
}
