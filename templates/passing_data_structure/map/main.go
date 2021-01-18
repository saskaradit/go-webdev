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
	// cars := []string{"Mercedes-Benz", "BMW", "Porsche"}
	cars := map[string]string{
		"Germany": "Mercedes-Benz",
		"Italy":   "Ferrari",
		"England": "Mclaren",
	}
	err := tpl.Execute(os.Stdout, cars)
	if err != nil {
		log.Fatalln(err)
	}
}
