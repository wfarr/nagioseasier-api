package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}
