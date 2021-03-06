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
		URL         *url.URL
		Submissions map[string][]string
		Header      http.Header
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
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
