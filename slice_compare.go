package main

import "fmt"

func compute(arr, brr []int) []int {
	x := []int{0, 0}
	for i := 0; i < len(arr); i++ {
		if arr[i] > brr[i] {
			x[0]++
		} else if arr[i] < brr[i] {
			x[1]++
		}
	}
	return x

}

func main() {
	arr := make([]int, 3)
	brr := make([]int, 3)
	fmt.Println("Enter 1st array :")
	for i := 0; i < len(arr); i++ {
		var ele int
		fmt.Println("Enter", i, "element:")
		fmt.Scanln(&ele)
		if ele > 100 || ele < 1 {
			fmt.Println("Wrong input")
			return
		}
		arr[i] = ele
	}
	fmt.Println("Enter 2nd array :")
	for i := 0; i < len(brr); i++ {
		var ele int
		fmt.Println("Enter", i, "element:")
		fmt.Scanln(&ele)
		if ele > 100 || ele < 1 {
			fmt.Println("Wrong input")
			return
		}
		brr[i] = ele
	}

	x := compute(arr, brr)
	fmt.Println(x)

}
