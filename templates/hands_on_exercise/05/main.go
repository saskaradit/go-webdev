package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	// parse csv
	records := prs("table.csv")

	// parse template
	tpl, err := template.ParseFiles("hw.go.html")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute Template
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func prs(filePath string) []Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	reader := csv.NewReader(src)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	return records
}
