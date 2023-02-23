package main

import (
	"fmt"
	"strings"
)

func capiFirstLast(arr []string) []string {
	brr := make([]string, 0)
	for _, v := range arr {
		k := len(v) - 1
		first := strings.ToUpper(string(v[0]))
		last := strings.ToUpper(string(v[k]))
		middle := v[1:k]
		res := first + middle + last
		brr = append(brr, res)
	}
	return brr
}

func main() {
	arr := []string{"nagarjun", "kumar", "arjun", "raju"}
	fmt.Println(capiFirstLast(arr))
}
