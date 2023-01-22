package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// creating the router
	router := mux.NewRouter()
	s := router.PathPrefix("/user").Subrouter()

	// post route
	s.HandleFunc("/create", CreateUser).Methods("POST")
	// get route
	s.HandleFunc("/getUser/{userId}", GetUserDetailsById).Methods("GET")
	// update route
	s.HandleFunc("/updateUser", UpdateUser).Methods("PUT")
	// delete route
	s.HandleFunc("/deleteUser/{userId}", DeleteUser).Methods("DELETE")
	http.ListenAndServe(":8000", s)

}
