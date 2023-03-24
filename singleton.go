package main

import (
	"fmt"
	"sync"
	"time"
)

type government struct {
	country string
	party   string
}

var instance *government

func getInstance(m *sync.Mutex) *government {
	m.Lock()
	defer m.Unlock()
	if instance == nil {
		fmt.Println("Creating Inststance ")
		instance = &government{"India", "XYZ"}
	} else {
		fmt.Println("Instance already created")
	}
	return instance
}

func main() {
	var m sync.Mutex
	for i := 0; i < 10; i++ {
		go getInstance(&m)
	}
	time.Sleep(2 * time.Second)
}
