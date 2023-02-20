package main

import (
	"fmt"
	"strings"
)

func palindromeCheck(str string) bool {
	bytes := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		bytes = append(bytes, str[i])
	}
	str2 := string(bytes)
	return strings.EqualFold(str, str2)
}

func main() {
	str := "hello"
	fmt.Println(palindromeCheck(str))

}
