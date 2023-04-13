package main

import "fmt"

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j = j - 1
		}
	}
	return arr
}

func main() {
	arr := []int{45, 23, 17, 38, 31}
	fmt.Println("Unsorted array:", arr)
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
