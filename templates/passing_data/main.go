package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", "Saskara")
	if err != nil {
		log.Fatalln(err)
	}
}
