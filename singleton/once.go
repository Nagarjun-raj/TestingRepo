package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var instance *single

func getInstance() *single {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating Single Instance Now")
				instance = &single{}
			})
	} else {
		fmt.Println("Single Instance already created")
	}
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
