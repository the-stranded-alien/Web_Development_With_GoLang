package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code you want in this function")
}

func main() {

}
