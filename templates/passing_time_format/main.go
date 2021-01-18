package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdateDMY": monthDayYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))

}

func monthDayYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
