package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", tesla)
	http.ListenAndServe(":8080", nil)
}

func tesla(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
	<!--Image does not serve-->
	<img src="/toby.jpg">
	`)
}
