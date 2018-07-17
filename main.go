package main

import (
	"log"
	"net/http"

	"./routes"
	"github.com/gorilla/mux"
)

// This API will provide a GET users endpoint
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/", routes.HandleGetOneUser).Methods("GET")
	router.HandleFunc("/api/users/", routes.HandleGetUsers).Methods("GET")
	router.HandleFunc("/api/", routes.HandleGetRoot)
	log.Fatal(http.ListenAndServe(":8082", router))
}
