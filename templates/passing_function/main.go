package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type car struct {
	Name string
	Make string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))

	// Does not work
	// tpl = template.Must(template.ParseFiles("index.html"))
	// tpl = tpl.Funcs(fm)
}
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	cayenne := car{"Cayenne", "Porsche"}
	gt := car{"GT63", "AMG"}
	m8 := car{"M8", "BMW"}
	cars := []car{cayenne, gt, m8}
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", cars)
	if err != nil {
		log.Fatalln(err)
	}
}
