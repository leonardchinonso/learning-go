package ch1

import (
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", handler) // any request will call "handler"
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		return 
	}
}