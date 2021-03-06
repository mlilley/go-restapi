package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	thing "github.com/mlilley/go-restapi/thing"
)

type App struct {
	router    *mux.Router
	thingRepo thing.ThingRepo
}

func NewApp() *App {
	thingRepo, err := thing.NewThingSqlite3Repo()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/things", thing.HandleList(thingRepo)).Methods("GET")
	router.HandleFunc("/things", thing.HandleCreate(thingRepo)).Methods("POST")
	router.HandleFunc("/things/{id}", thing.HandleGet(thingRepo)).Methods("GET")
	router.HandleFunc("/things/{id}", thing.HandleUpdate(thingRepo)).Methods("PUT")
	router.HandleFunc("/things/{id}", thing.HandleDelete(thingRepo)).Methods("DELETE")

	return &App{router: router, thingRepo: thingRepo}
}

func main() {
	app := NewApp()
	log.Fatal(http.ListenAndServe(":8000", app.router))
}
