package handlers

import (
	"fmt"
	"net/http"
	"simple_crud/internal/storage"
	"strconv"
)

func ReadIdHandler(dbStorage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/tasks/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task, err := dbStorage.ReadId(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		answer := fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nStatus: %s\nCreate at: %s\nLast update: %s\n",
			task.Id, task.Title, task.Description, task.Status, task.CreatedAt.String(), task.LastUpdate.String(),
		)

		w.Write([]byte(answer))
	}
}
