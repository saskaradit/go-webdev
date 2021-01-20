package main

import (
	"io"
	"net/http"
)

type benz int

func (b benz) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "The best or nothing")
}

type amg int

func (a amg) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "AMG GT 4 Door")
}

func main() {
	var b benz
	var a amg
	mux := http.NewServeMux()
	mux.Handle("/car/", b)
	mux.Handle("/amg", a)
	http.ListenAndServe(":8080", mux)
}
