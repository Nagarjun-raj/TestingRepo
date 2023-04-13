package main

import "fmt"

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

func main() {
	arr := []int{4, 5, 2, 7, 1, 6, 3}
	fmt.Println("Unsorted array:", arr)
	selectionSort(arr)
	fmt.Println("Sorted array:", arr)
}
