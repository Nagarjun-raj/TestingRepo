package main

import "fmt"

func factorialCalc(num int) int {
	res := 1
	for ; num != 0; num-- {
		res *= num
	}
	return res
}

func main() {
	num := 2
	fmt.Println(factorialCalc(num))

}
