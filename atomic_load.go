package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func hi(i *int32) {
	atomic.AddInt32(i, 1)
}

func main() {
	var i int32
	for j := 0; j < 10; j++ {
		go hi(&i)
	}
	time.Sleep(time.Second)
	fmt.Println("i value: ", atomic.LoadInt32(&i))

}
