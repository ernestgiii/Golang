// This script creates a simple webserver for Go
package main

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 color=forestgreen>BlackNTech Media</h1>"))
}

func main() {
	http.HandleFunc("/hello", Hello)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
