package main

import (
	"github.com/gorilla/mux"
	"log"
	"app/controllers"
	"net/http"
)

func main() {
	log.Print("Starting main")
	// Controllers declaration
	c := &controllers.UserController{}
	i := &controllers.IndexController{}
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/v1/").Subrouter()

	// Routes handling
	r.HandleFunc("/", i.Index)
	apiRouter.HandleFunc("/users", c.UserGetAll).Methods("GET")
	apiRouter.HandleFunc("/users/add", c.UserAdd).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
