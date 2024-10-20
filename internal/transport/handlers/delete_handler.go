package handlers

import (
	"net/http"
	"simple_crud/internal/storage"
	"strconv"
)

func DeleteHandler(dbStorage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.URL.Path[len("/tasks/"):]

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = dbStorage.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
