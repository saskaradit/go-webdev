package main

import (
	"fmt"
	"net/http"
)

type tesla int

func (t tesla) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "you can put any code you want")
}

func main() {
	var t tesla
	http.ListenAndServe(":8080", t)
}
