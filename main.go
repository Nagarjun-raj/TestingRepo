package main

import "fmt"

func find(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max

}

func main() {
	arr := []int{10, 5, 3, 18, 9}
	fmt.Println(find(arr))

}
