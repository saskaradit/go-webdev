package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type car struct {
	Name string
	Make string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	cayenne := car{"Cayenne", "Porsche"}
	gt := car{"GT63", "AMG"}
	m8 := car{"M8", "BMW"}
	cars := []car{cayenne, gt, m8}
	err := tpl.Execute(os.Stdout, cars)
	if err != nil {
		log.Fatalln(err)
	}
}
