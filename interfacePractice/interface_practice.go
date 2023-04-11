package main

import "fmt"

type calculate1 interface {
	area() int
}

type square struct {
	length int
}

func (s square) area() int {
	return s.length * s.length
}

type cube struct {
	edge int
}

type calculate2 interface {
	calculate1
	volume() int
}

func (c cube) area() int {
	return 6 * c.edge * c.edge
}

func (c cube) volume() int {
	return c.edge * c.edge * c.edge
}

func totalArea(s []calculate1) int {
	total := 0
	for _, v := range s {
		total += v.area()
	}
	return total

}

func totalVolume(c cube) int {
	total := c.volume()
	return total

}

func main() {
	s := square{6}
	c := cube{6}
	arr := []calculate1{s, c}
	totalArea := totalArea(arr)
	totalVolume := totalVolume(c)
	fmt.Println("Total area :", totalArea)
	fmt.Println("Total Volme:", totalVolume)

}
