package main

import "fmt"

func reverse(str string) string {
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}

func main() {
	str := "Slice"
	fmt.Println(reverse(str))
}
