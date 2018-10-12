package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type thing struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Parts []part `json:"parts,omitempty"`
}

type part struct {
	Name string `json:"name,omitempty"`
}

var things []thing

func getThings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(things)
}

func createThing(w http.ResponseWriter, r *http.Request) {
	var t thing
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = string(len(things) + 1)
	things = append(things, t)
	json.NewEncoder(w).Encode(things)
}

func getThing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, t := range things {
		if t.ID == params["id"] {
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	w.WriteHeader(404)
}

func updateThing(w http.ResponseWriter, r *http.Request) {
	var new thing
	json.NewDecoder(r.Body).Decode(&new)

	params := mux.Vars(r)
	for i, old := range things {
		if old.ID == params["id"] {
			new.ID = old.ID
			things[i] = new
			json.NewEncoder(w).Encode(new)
			return
		}
	}
	w.WriteHeader(404)
}

func deleteThing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, old := range things {
		if old.ID == params["id"] {
			things = append(things[:i], things[i+1:]...)
			w.WriteHeader(204)
			return
		}
	}
}

func main() {
	router := mux.NewRouter()

	things = append(things, thing{ID: "1", Name: "One", Parts: []part{part{Name: "B1"}, part{Name: "B2"}}})
	things = append(things, thing{ID: "2", Name: "Two", Parts: []part{part{Name: "B3"}, part{Name: "B4"}}})

	router.HandleFunc("/things", getThings).Methods("GET")
	router.HandleFunc("/things", createThing).Methods("POST")
	router.HandleFunc("/things/{id}", getThing).Methods("GET")
	router.HandleFunc("/things/{id}", updateThing).Methods("PUT")
	router.HandleFunc("/things/{id}", deleteThing).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
