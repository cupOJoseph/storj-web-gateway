package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "3000", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/ipfs", handleRoute)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request %r", r.URL.Path)
}
