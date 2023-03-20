package main

import "fmt"

func calculateRatio(arr []int) (float32, float32, float32, float32) {
	size := len(arr)
	var pos, neg, zer int
	for _, v := range arr {
		if v > 0 {
			pos++
		} else if v < 0 {
			neg++
		} else {
			zer++
		}
	}
	sizef := float32(size)
	posf := float32(pos)
	negf := float32(neg)
	zerf := float32(zer)
	return sizef, posf, negf, zerf
}

func main() {
	arr := []int{-8, 0, 2, 4, -9}
	sizef, posf, negf, zerf := calculateRatio(arr)
	fmt.Println("Ratios are : ")
	fmt.Println("Positive nums: ", posf/sizef)
	fmt.Println("Negative nums: ", negf/sizef)
	fmt.Println("Zeros : ", zerf/sizef)
}
