package main

import (
	"fmt"
	"net/http"
	"time"
)

func execution(h http.HandlerFunc) http.HandlerFunc {
	a := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		end := time.Now()
		fmt.Printf("Time taken: %v\n", end.Sub(start))

	}
	return a
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Nagarjun!")
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Prashanth!")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to Decorator pattern")
}

func main() {
	http.HandleFunc("/hello", execution(hello))
	http.HandleFunc("/hi", execution(hi))
	http.HandleFunc("/welcome", execution(welcome))
	http.ListenAndServe(":8123", nil)
}
