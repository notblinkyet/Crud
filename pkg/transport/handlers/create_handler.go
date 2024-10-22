package handlers

import (
	"encoding/json"
	"net/http"
	"simple_crud/pkg/models"
	"simple_crud/pkg/storage"
	"strconv"
)

func CreateHandler(dbStorage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Need POST method\n", http.StatusBadRequest)
		}
		var task models.Task
		var createMess string = "Create new task with id "

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := dbStorage.Create(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(createMess))
		w.Write([]byte(strconv.Itoa(id)))
		w.Write([]byte("\n"))
	}
}
