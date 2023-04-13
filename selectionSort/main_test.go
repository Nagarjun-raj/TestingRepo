package main

import "testing"

func BenchmarkSelection(b *testing.B) {
	arr := []int{4, 2, 1, 9, 14, 10}
	for i := 0; i < b.N; i++ {
		selection(arr)

	}
}
