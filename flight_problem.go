package main

import (
	"errors"
	"fmt"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []Flight) (int, int, error) {
	var err error
	for i := 0; i < len(flights); i++ {
		for j := i + 1; j < len(flights); j++ {
			if flights[i].Price > flights[j].Price {
				flights[i], flights[j] = flights[j], flights[i]
			}
			if flights[i].Price <= 0 {
				err = errors.New("Invalid Price")
			}
		}
	}
	l := len(flights) - 1
	return flights[0].Price, flights[l].Price, err

}

func main() {
	hello := []Flight{
		Flight{"India", "USA", 40000},
		Flight{"India", "Germany", 25000},
		Flight{"India", "Berlin", 60000},
		Flight{"India", "Dubai", -8067},
	}
	fmt.Println(GetMinMax(hello))
}
