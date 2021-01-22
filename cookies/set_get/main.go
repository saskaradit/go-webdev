package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "Rad-Cookie",
		Value: "Jengjet",
	})
	fmt.Fprintln(res, "Cookie written check your browser")
}

func read(res http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("Rad-Cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "Your cookie :", c1)
	}

	c2, err := req.Cookie("General")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "Your cookie :", c2)
	}

	c3, err := req.Cookie("Specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "Your cookie :", c3)
	}
}

func abundance(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "General",
		Value: "Jengjet",
	})
	http.SetCookie(res, &http.Cookie{
		Name:  "Specific",
		Value: "Jengjet",
	})

	fmt.Fprintln(res, "Cookies Written ")
}
