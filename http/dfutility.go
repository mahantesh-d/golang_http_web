package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe("10.138.32.75:22", nil))
}
