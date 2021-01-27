package main

import (
	"Udemy/GoWebDev/mongodb/06_hands-on/starting-code/controllers"
	"Udemy/GoWebDev/mongodb/06_hands-on/starting-code/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return make(map[string]models.User)
}
