package main

import (
	"io"
	"net/http"
)

type benz int

func (b benz) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/car":
		io.WriteString(res, "AMG GT 4 Door")
	case "/about":
		io.WriteString(res, "The best or nothing")
	}
}

func main() {
	var b benz
	http.ListenAndServe(":8080", b)
}
