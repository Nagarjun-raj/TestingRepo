package people

import (
	"fmt"
	"sync"
)

type people struct {
	nose int
	ear  int
}

var instance *people

func GetInstance(once *sync.Once) *people {
	if instance == nil {
		once.Do(func() {
			fmt.Println("Creating instance:")
			instance = &people{1, 2}
		})
	} else {
		fmt.Println("Instance already created")
	}
	return instance
}
