package main

import (
    "net/http"
    "fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (input String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, input)
}

func (input Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, input.Greeting + " " + input.Punct + " " + input.Who)
}

func main() {
    http.Handle("/string", String("I'm a frayed knot."))
	  http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}

