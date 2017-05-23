package main

import (
	"github.com/gorilla/mux"
	"log"
	"myapp/controllers"
	"net/http"
)

func main() {
	log.Print("Starting main")

	// Controllers declaration
	c := &controllers.UserController{}
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	// Routes handling
	s.HandleFunc("/users", c.GetUser).Methods("GET")
	// s.HandleFunc("/users/add", c.AddUser).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
