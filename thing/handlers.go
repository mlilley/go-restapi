package thing

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleList(repo *Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		ts := repo.GetAll()

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(ts); err != nil {
			panic(err)
		}
	}
}

func HandleCreate(repo *Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		var t Thing

		if err := json.NewDecoder(io.LimitReader(r.Body, 1048576)).Decode(&t); err != nil {
			http.Error(w, "invalid json", http.StatusUnprocessableEntity)
			return
		}

		tt := repo.Create(t)

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(*tt); err != nil {
			panic(err)
		}
	}
}

func HandleGet(repo *Repo) http.HandlerFunc {
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

		id, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		t := repo.Get(id)
		if t == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(*t); err != nil {
			panic(err)
		}
	}
}

func HandleUpdate(repo *Repo) http.HandlerFunc {
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

		id, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		var t Thing

		if err := json.NewDecoder(io.LimitReader(r.Body, 1048576)).Decode(&t); err != nil {
			http.Error(w, "invalid json", http.StatusUnprocessableEntity)
			return
		}

		tt := repo.Update(id, t)
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

func HandleDelete(repo *Repo) http.HandlerFunc {
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

		id, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		if ok := repo.Delete(id); !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func setJSONContentType(w http.ResponseWriter) {

}
