package main

import (
	
	"github.com/gorilla/mux"
	"log"
	"net/http"
	
)

func handleRequest() {

	const port string = ":8000"
	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/person", addPerson).Methods("POST")
	router.HandleFunc("/person", editPerson).Methods("PUT")
	
	log.Println("Server Listening on Port",port)
	log.Fatalln(http.ListenAndServe(port, router))
}

func main() {
	handleRequest()
}



