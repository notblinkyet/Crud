package handlers

import (
	"fmt"
	"net/http"

	"github.com/notblinkyet/Crud/pkg/storage"
)

func ReadAllHandler(dbStorage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Need GET method\n", http.StatusBadRequest)
			return
		}
		tasks, err := dbStorage.ReadAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write([]byte("All data:\n"))
		for _, task := range tasks {
			answer := fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nStatus: %s\nCreate at: %s\nLast update: %s\n\n",
				task.Id, task.Title, task.Description, task.Status, task.CreatedAt.String(), task.LastUpdate.String(),
			)

			w.Write([]byte(answer))
		}
	}
}
