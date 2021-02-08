package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello root, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
