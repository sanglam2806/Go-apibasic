package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Persons []Person `json:"person"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname`
	LastName  string `json:"lastname"`
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	var response Response

	persons := preparePersons()

	response.Persons = persons

	// update content-type
	w.Header().Set("Content-Type", "application/json")

	// specify HTTP status code
	w.WriteHeader(http.StatusOK)

	// convert Struts to json
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	// update response
	w.Write(jsonResponse)

}

func preparePersons() []Person {
	var persons []Person

	var person Person

	person.Id = 1
	person.FirstName = "Lam"
	person.LastName = "Sang"

	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Tim"
	person.LastName = "Mitsukeru"

	persons = append(persons, person)

	return persons
}

func StartWebServer() error {
	fmt.Println("Rest API with mux Routers")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rootPage).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")

	http.ListenAndServe(":8080", router)

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
