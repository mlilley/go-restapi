package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	thing "github.com/mlilley/go-restapi/thing"
)

type Server struct {
	router *mux.Router
}

func NewServer() Server {
	thingRepo := thing.NewRepo()
	r := mux.NewRouter()
	s := Server{router: r}

	addRoute(r, "/things", "GET", thing.HandleList(thingRepo))
	addRoute(r, "/things", "POST", thing.HandleCreate(thingRepo))
	addRoute(r, "/things/{id}", "GET", thing.HandleGet(thingRepo))
	addRoute(r, "/things/{id}", "PUT", thing.HandleUpdate(thingRepo))
	addRoute(r, "/things/{id}", "DELETE", thing.HandleDelete(thingRepo))

	return s
}

func addRoute(r *mux.Router, path string, methods string, handler http.HandlerFunc) {
	r.HandleFunc(path, handler).Methods(methods)
}

func main() {
	s := NewServer()
	log.Fatal(http.ListenAndServe(":8000", s.router))
}
