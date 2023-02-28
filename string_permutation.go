package main

import "fmt"

func CheckPermutations(str1, str2 string) bool {
	temp := make(map[rune]int)
	if len(str1) != len(str2) {
		return false
	}
	for _, rune := range str1 {
		temp[rune] += 1
	}
	for _, rune := range str2 {
		temp[rune] -= 1
		if temp[rune] < 0 {
			return false
		}
	}
	return true
}

func main() {
	str1 := "adcme"
	str2 := "medak"
	isPermutation := CheckPermutations(str1, str2)
	fmt.Println(isPermutation)

}
