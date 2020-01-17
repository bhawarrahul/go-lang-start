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

	people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy"})
	people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
	http.HandleFunc("/people", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
