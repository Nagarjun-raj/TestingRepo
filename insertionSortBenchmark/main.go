package main

import "fmt"

func insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j:=i-1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{6, 9, 2, 7, 10, 14, 12}
	brr := insertion(arr)
	fmt.Println(brr)
}
