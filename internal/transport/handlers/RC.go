package handlers

import (
	"net/http"
	"simple_crud/internal/storage"
)

func RCHandler(dbStorage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ReadAllHandler(dbStorage)(w, r)
		case http.MethodPost:
			CreateHandler(dbStorage)(w, r)
		default:
			http.Error(w, "Need method: DELETE, GET, PUT or PATCH", http.StatusBadRequest)
			return
		}
	}
}
