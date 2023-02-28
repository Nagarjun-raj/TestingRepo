package main

import "fmt"

func search(arr []int) []int {
	var brr []int
	for _, v := range arr {
		if v%5 == 0 {
			brr = append(brr, v)
		}
	}
	return brr
}

func main() {
	arr := []int{12, 34, 90, 27, 65}
	fmt.Println(search(arr))
}
