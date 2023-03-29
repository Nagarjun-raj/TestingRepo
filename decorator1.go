package main

import "fmt"

type hello interface {
	hi() string
}

type basic struct {
}

func (b *basic) hi() string {
	return "Hello"
}

type decorator1 struct {
	h hello
}

func (d *decorator1) hi() string {
	return d.h.hi() + " Nagarjun"
}

type decorator2 struct {
	h hello
}

func (d *decorator2) hi() string {
	return d.h.hi() + " How are you"
}

func main() {
	basic := &basic{}
	decorator1 := &decorator1{basic}
	decorator2 := &decorator2{decorator1}
	fmt.Println(decorator2.hi())

}
