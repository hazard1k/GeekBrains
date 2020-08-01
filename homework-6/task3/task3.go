package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.URL.Query() {
		fmt.Fprint(w, key+" = "+strings.Join(value, ",")+"\n")
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", handlerHello)
	http.ListenAndServe(":80", nil)
}
