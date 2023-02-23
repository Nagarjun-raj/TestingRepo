package main

import "fmt"

func countFirstLetter(names []string) map[string]int {
	frequency := make(map[string]int)
	for _, v := range names {
		firstLetter := string(v[0])
		frequency[firstLetter]++
	}
	return frequency
}

func main() {
	names := []string{"Roy", "Ted", "Nick", "Chris", "Tony"}
	frequency := countFirstLetter(names)
	fmt.Println(frequency)
}
