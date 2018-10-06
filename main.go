package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Superhero struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Superpower string `json:"superpower,omitempty"`
}

var superheroes []Superhero

func main() {
	superheroes = append(superheroes, Superhero{ID: "-1", Name: "MissingNo.", Superpower: "Item duplication"})
	superheroes = append(superheroes, Superhero{ID: "1", Name: "Dodger", Superpower: "Dodger of all things unwanted"})

	router := mux.NewRouter()
	router.HandleFunc("/superhero", GetSuperheroes).Methods("GET")
	router.HandleFunc("/superhero/{id}", GetSuperhero).Methods("GET")
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
