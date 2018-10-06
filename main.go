package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

type Superhero struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Superpower string `json:"superpower,omitempty"`
}

var superheroes []Superhero

func main() {
	createRecords()
	createRoutes()
}

func createRecords() {
	superheroes = append(superheroes, Superhero{ID: "-1", Name: "MissingNo.", Superpower: "Item duplication"})
	superheroes = append(superheroes, Superhero{ID: "1", Name: "Dodger", Superpower: "Dodger of all things unwanted"})
	superheroes = append(superheroes, Superhero{ID: "2", Name: "Freezer", Superpower: "Makes things really, really cold"})
	superheroes = append(superheroes, Superhero{ID: "3", Name: "Mr. Owl", Superpower: "Knows how many licks it takes"})
}

func createRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/superhero", GetSuperheroes).Methods("GET")
	router.HandleFunc("/superhero/{id}", GetSuperhero).Methods("GET")
	router.HandleFunc("/superhero/{id}", PutSuperhero).Methods("PUT")
	router.HandleFunc("/superhero/{id}", PostSuperhero).Methods("POST")
	router.HandleFunc("/superhero/{id}", DeleteSuperhero).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetSuperhero(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for _, hero := range superheroes {
		if hero.ID == params["id"] {
			json.NewEncoder(writer).Encode(hero)
			return
		}
	}
	json.NewEncoder(writer).Encode(superheroes[0])
}

func GetSuperheroes(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(superheroes)
}

func PutSuperhero(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var hero Superhero
	_ = json.NewDecoder(request.Body).Decode(&hero)
	hero.ID = params["id"]
	superheroes = append(superheroes, hero)
	json.NewEncoder(writer).Encode(hero)
}

func PostSuperhero(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var hero Superhero
	_ = json.NewDecoder(request.Body).Decode(&hero)
	hero.ID = params["id"]
	superheroes = append(superheroes, hero)
	json.NewEncoder(writer).Encode(hero)
}

func DeleteSuperhero(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, hero := range superheroes {
		if hero.ID == params["id"] {
			superheroes = append(superheroes[:index], superheroes[index+1:]...)
			json.NewEncoder(writer).Encode(hero)
			break
		}
	}
}
