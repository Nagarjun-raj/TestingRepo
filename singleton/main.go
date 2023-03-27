package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type single struct {
}

var instance *single

func getInstance() *single {
	// mutex.Lock()
	// defer mutex.Unlock()
	// fmt.Println("hello")
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance")
			instance = &single{}
		} else {
			fmt.Println("Single instance already created-1")
		}
	} else {
		fmt.Println("Single instance already created-2")
	}
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
