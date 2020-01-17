package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var people []Person

func main() {

	people = append(people, Person{ID: "315", Firstname: "Rahul", Lastname: "Bhawar"})
	people = append(people, Person{ID: "316", Firstname: "Amol", Lastname: "Edewar"})
	http.HandleFunc("/people", indexHandler)
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to GoLang API call - http://localhost:8080/people to get pelope List")
}
