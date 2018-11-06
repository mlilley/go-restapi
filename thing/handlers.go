package thing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleList(repo ThingRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		ts, err := repo.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(ts); err != nil {
			panic(err)
		}
	}
}

func HandleCreate(repo ThingRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		var t Thing
		if err := json.NewDecoder(io.LimitReader(r.Body, 1048576)).Decode(&t); err != nil {
			http.Error(w, "invalid json", http.StatusUnprocessableEntity)
			return
		}

		tt, err := repo.Create(&t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("%v\n", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(*tt); err != nil {
			panic(err)
		}
	}
}

func HandleGet(repo ThingRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		params := mux.Vars(r)

		strID, ok := params["id"]
		if !ok {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		t, err := repo.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if t == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(*t); err != nil {
			panic(err)
		}
	}
}

func HandleUpdate(repo ThingRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		params := mux.Vars(r)
		if params == nil {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}

		strID, ok := params["id"]
		if !ok {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		var t Thing
		if err := json.NewDecoder(io.LimitReader(r.Body, 1048576)).Decode(&t); err != nil {
			http.Error(w, "invalid json", http.StatusUnprocessableEntity)
			return
		}

		tt, err := repo.Update(id, &t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if tt == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(*tt); err != nil {
			panic(err)
		}
	}
}

func HandleDelete(repo ThingRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		params := mux.Vars(r)
		if params == nil {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}

		strID, ok := params["id"]
		if !ok {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		deleted, err := repo.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !deleted {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
