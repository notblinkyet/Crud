package handlers

import (
	"encoding/json"
	"net/http"
	"simple_crud/internal/storage"
	"strconv"
)

type Update struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Statis      string `json:"status"`
}

func UpdateHander(dbStorage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var update Update

		idStr := r.URL.Path[len("/tasks/"):]

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		json.NewDecoder(r.Body).Decode(&update)

		err = dbStorage.Update(id, update.Title, update.Description, update.Statis)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
