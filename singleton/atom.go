package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func atomicLoad(x *int32) {
	atomic.AddInt32(x, 1)
}
func main() {
	var x int32
	for i := 0; i < 10; i++ {
		go atomicLoad(&x)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("value of x:", atomic.LoadInt32(&x))
}
