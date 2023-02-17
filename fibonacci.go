package main

import "fmt"

func main() {
	a, b := 0, 1
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		if a%2 == 0 {
			sum += a
		}
		temp := a + b
		a = b
		b = temp
	}
	fmt.Printf("Sum:%d\n", sum)
}
