package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var x int32 = 0

func atomicLoad() {
	atomic.AddInt32(&x, 1)
}
func main() {
	for i := 0; i < 10; i++ {
		go atomicLoad()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("value of x:", atomic.LoadInt32(&x))
}
