package main

import (
	"fmt"
	"net/http"
)

var counter int = 0

// Ignore favicon requests
func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "You are visitor number %d\n", counter)

	subPath := r.URL.Path[1:2]
	fmt.Fprintf(w, "First character of path is: %s\n", subPath)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
