package main

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"golang-firestore/repository"
	"golang-firestore/entity"
)



var (
repo repository.PostRepository = repository.NewPostRepository() 
)


func getPeople(response http.ResponseWriter , request *http.Request)  {
	response.Header().Set("Content-Type","application/json")
	people , err := repo.FindAll()
	if err !=nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error Getting the Posts"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(people)
}

func addPerson(response http.ResponseWriter , request *http.Request){
	response.Header().Set("Content-Type","application/json")
	var Person entity.Person
	err:= json.NewDecoder(request.Body).Decode(&Person)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error Unmarshalling Data"}`))
		return
	}
	Person.ID = rand.Int63()
	repo.AddPersonRepo(&Person)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(Person)
}

func editPerson(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type","application/json")
	var Person entity.Person
	err:= json.NewDecoder(request.Body).Decode(&Person)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error Unmarshalling Data"}`))
		return
	}
	repo.EditPersonRepo(&Person)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(Person)
}