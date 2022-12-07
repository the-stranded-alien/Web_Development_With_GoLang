package main

import (
	"io"
	"net/http"
)

type hotdog int

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

type hotcat int

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8888", nil)
}
