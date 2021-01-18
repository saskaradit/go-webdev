package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.go.html"))
}

func main() {
	hotels := region{
		Region: "Southern",
		Hotels: []hotel{
			hotel{
				"Rad", "Groove Street", "California", "9090",
			},
			hotel{
				"Jengjet", "Brit Street", "California", "9090",
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
