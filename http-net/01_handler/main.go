package main

import (
	"fmt"
	"net/http"
)

type anything int

// implicitly implements the Handler interface

func (a anything) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
