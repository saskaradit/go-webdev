package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", chin)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.go.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(res, "dog.go.html", nil)
}

func chin(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "dog.jpg")
}
