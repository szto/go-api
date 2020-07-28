package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	u := String("Hellow World")
	http.Handle("/", h)
	http.HandleFunc("/string", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, u)
	})
	http.HandleFunc("/struct", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, &Struct{"Hello", ":", "Gphoers!"})
	})
	// http.Handle("/struct", )
	http.ListenAndServe("localhost:4000", nil)
}
