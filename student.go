package main

import (
	"fmt"
)

type student struct {
	rollno int
	name   string
}

func sorting(arr []student, ch int) []student {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if ch == 2 {
				if arr[i].rollno < arr[j].rollno {
					arr[i], arr[j] = arr[j], arr[i]
				}
			} else if ch == 1 {
				if arr[i].rollno > arr[j].rollno {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}
	return arr
}

func main() {
	var choice int
	s1 := student{23, "Rahul"}
	s2 := student{88, "Nagarjun"}
	s3 := student{10, "Shivu"}
	s4 := student{45, "Vaibhav"}
	s5 := student{18, "Virat"}
	s := []student{s1, s2, s3, s4, s5}
	fmt.Println(s)
	fmt.Println("Enter 1 for ascending or 2 for descending order")
	fmt.Scanln(&choice)
	fmt.Println(sorting(s, choice))
}

// sort.Slice(s, func(i, j int) bool {
// 	return s[i].rollno < s[j].rollno
// })
// fmt.Println(s)
