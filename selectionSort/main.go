package main

import "fmt"

func selection(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		//Finding index of lowest element
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}

func main() {
	arr := []int{4, 2, 1, 9, 14, 10}
	x := selection(arr)
	fmt.Println(x)
}
