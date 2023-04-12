package main

import "fmt"

func binary(arr []int, num int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if num > arr[mid] {
			low = mid + 1
		} else if num == arr[mid] {
			return mid
		} else {
			high = mid - 1
		}

	}
	return -1

}
func main() {
	arr := []int{2, 4, 6, 8, 10}
	num := 8
	x := binary(arr, num)
	if x == -1 {
		fmt.Println("Element not present in the slice")
	} else {
		fmt.Printf("%d is present at the index %d,", num, x)
	}
}
