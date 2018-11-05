package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	thing "github.com/mlilley/go-restapi/mutex/thing"
)

type App struct {
	router    *mux.Router
	thingRepo *thing.Repo
}

func NewApp() *App {
	thingRepo := thing.NewRepo()

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
