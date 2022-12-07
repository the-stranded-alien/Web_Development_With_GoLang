package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))
	http.ListenAndServe(":8888", nil)
}
