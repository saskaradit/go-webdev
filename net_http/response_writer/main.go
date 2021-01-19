package main

import (
	"fmt"
	"net/http"
)

type tesla int

func (t tesla) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Rad-Key", "This is rad's Server")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h1>You put code in here</h1>")
}

func main() {
	var t tesla
	http.ListenAndServe(":8080", t)
}
