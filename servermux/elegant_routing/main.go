package main

import (
	"io"
	"net/http"
)

func benz(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "The best or nothing")
}

func amg(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "AMG GT 4 Door")
}

func main() {
	http.HandleFunc("/car/", benz)
	http.HandleFunc("/amg", amg)
	http.ListenAndServe(":8080", nil)
}
