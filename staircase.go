package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the size of staircase: ")
	fmt.Scanln(&size)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}
