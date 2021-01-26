package main

import (
	"html/template"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	c = appendValue(w, c)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

// getCookie function returns a cookie
func getCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(res, cookie)
	}
	return cookie
}

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// Value
	pic1 := "jengjet.jpg"
	pic2 := "rad.jpg"
	pic3 := "benz.jpg"

	// Append
	s := c.Value
	if !strings.Contains(s, pic1) {
		s += "|" + pic1
	}
	if !strings.Contains(s, pic2) {
		s += "|" + pic2
	}
	if !strings.Contains(s, pic3) {
		s += "|" + pic3
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
