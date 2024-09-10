package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var people []Person

func main() {
	fmt.Print(1)
	fmt.Print(2)
	fmt.Print("new Branch KAtkov3")
	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/health", healthHandler)
	log.Println("listening 8080...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPost:
		postPerson(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
	fmt.Fprintf(w, "get people: '%v", people)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	people = append(people, person)
	fmt.Fprintf(w, "post new person '%v", person)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "http server works correctly")
}
