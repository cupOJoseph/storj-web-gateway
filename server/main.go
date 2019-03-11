package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received request for path: %s", r.URL.Path[1:])
}

func main() {
	port := strconv.Itoa(5000)
	fmt.Printf("Listening to traffic on port " + port + "\n")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
