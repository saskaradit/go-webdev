package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type tesla int

func (t tesla) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.html", data)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var t tesla
	http.ListenAndServe(":8080", t)
}
