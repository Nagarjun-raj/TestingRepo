package main

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{4, 2, 3, 1, 7, 3}
		bubbleSort(arr)
	}
}
