package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, client"))
}
