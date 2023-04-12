package main

import (
	"fmt"
)

func bubbleSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				// Swap the elements
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func main() {
	array := []int{4, 2, 3, 1, 7, 3}
	fmt.Println("Unsorted array:", array)
	bubbleSort(array)
	fmt.Println("Sorted array:", array)
}
