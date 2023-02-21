package main

import "fmt"

func ToUppercase(s string) string {
	var a []rune
	for _, rune := range s {
		if rune >= 'a' && rune <= 'z' {
			rune -= 32
		}
		a = append(a, rune)
	}
	return string(a)
}

func main() {
	name := "prasHanTh"
	fmt.Println(ToUppercase(name))
}
