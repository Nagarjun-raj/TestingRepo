package main

import "testing"

func BenchmarkInsertion(b *testing.B) {
	test := []int{6, 9, 2, 7, 10, 14, 12}
	for i := 0; i < b.N; i++ {
		insertion(test)
	}
}
