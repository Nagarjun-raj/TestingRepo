package main

import (
	"once/people"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		go people.GetInstance(&once)
	}
	time.Sleep(3 * time.Second)

}
