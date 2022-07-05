package main

import (
	"encoding/json"
	"fmt"
	"httpserver/entity"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

var person entity.Person

func main() {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {

		r.Get("/person", GetPerson)
		r.Post("/create", CreatePerson)
		r.Put("/update", UpdatePerson)
		//router.Delete("/delete",DeletePerson)
	})

	http.HandleFunc("/", handleTest1)
	http.ListenAndServe(":8080", router)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	name := r.URL.Query().Get("name")
	gender := r.URL.Query().Get("gender")

	name = fmt.Sprintf("Welcome to our internship, %s", name)
	gender = fmt.Sprintf(" gender %s", gender)

	fmt.Fprint(w, name, gender)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)

	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &person)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)

	var updatedPerson entity.Person

	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &updatedPerson)
	person.ID = updatedPerson.ID
	person.Name = updatedPerson.Name
	person.Gender = updatedPerson.Gender
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}

func handleTest1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "nimic")
}

// create 4 apis: create person, read person, delete person, update person
