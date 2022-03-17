package ch1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func Serve2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler2 echoes the Path of the requested URL
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		return
	}
}

// counter echoes the number of request calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, err := fmt.Fprintf(w, "Count %d\n", count)
	if err != nil {
		return
	}
	mu.Unlock()
}

