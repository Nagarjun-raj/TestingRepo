package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32

func hi() {
	atomic.AddInt32(&i, 1)
}

func main() {
	for j := 0; j < 10; j++ {
		go hi()
	}
	time.Sleep(time.Second)
	fmt.Println("i value: ", atomic.LoadInt32(&i))

}
