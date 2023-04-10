package main

import "fmt"

func doubleChars(original string) string {
	new := make([]rune, 0, len(original)*2)
	for _, v := range original {
		new = append(new, v, v)
	}
	return string(new)
}

func main() {
	original := "Channels"
	doubled := doubleChars(original)
	fmt.Println(doubled)
}
