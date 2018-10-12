package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: protect shared data with mutex etc

type server struct {
	router *mux.Router
}

type thing struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Parts []part `json:"parts,omitempty"`
}

type part struct {
	Name string `json:"name,omitempty"`
}

var things []thing

func (s *server) routes() {
	s.router.HandleFunc("/things", s.handleGetThings()).Methods("GET")
	s.router.HandleFunc("/things", s.handleCreateThing()).Methods("POST")
	s.router.HandleFunc("/things/{id}", s.handleGetThing()).Methods("GET")
	s.router.HandleFunc("/things/{id}", s.handleUpdateThing()).Methods("PUT")
	s.router.HandleFunc("/things/{id}", s.handleDeleteThing()).Methods("DELETE")
}

func (s *server) handleGetThings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(things)
	}
}

func (s *server) handleCreateThing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t thing
		json.NewDecoder(r.Body).Decode(&t)
		t.ID = string(len(things) + 1)
		things = append(things, t)
		json.NewEncoder(w).Encode(things)
	}
}

func (s *server) handleGetThing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		for _, t := range things {
			if t.ID == params["id"] {
				json.NewEncoder(w).Encode(t)
				return
			}
		}
		w.WriteHeader(404)
	}
}

func (s *server) handleUpdateThing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func (s *server) handleDeleteThing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		for i, old := range things {
			if old.ID == params["id"] {
				things = append(things[:i], things[i+1:]...)
				w.WriteHeader(204)
				return
			}
		}
		w.WriteHeader(404)
	}
}

func main() {
	s := server{router: mux.NewRouter()}
	s.routes()

	things = append(things, thing{ID: "1", Name: "One", Parts: []part{part{Name: "B1"}, part{Name: "B2"}}})
	things = append(things, thing{ID: "2", Name: "Two", Parts: []part{part{Name: "B3"}, part{Name: "B4"}}})

	log.Fatal(http.ListenAndServe(":8000", s.router))
}
