package main

import "fmt"

func checkLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	year := 2020
	if checkLeapYear(year) {
		fmt.Printf("%d is a Leap Year\n", year)
		return
	}
	fmt.Printf("%d is not a Leap Year\n", year)
}
