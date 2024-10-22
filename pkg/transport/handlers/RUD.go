package handlers

import (
	"net/http"
	"simple_crud/pkg/storage"
)

func RUDHandler(dbStorage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			DeleteHandler(dbStorage)(w, r)
		case http.MethodGet:
			ReadIdHandler(dbStorage)(w, r)
		case http.MethodPut:
			UpdateHander(dbStorage)(w, r)
		case http.MethodPatch:
			UpdateHander(dbStorage)(w, r)
		default:
			http.Error(w, "Need method: DELETE, GET, PUT or PATCH", http.StatusBadRequest)
			return
		}
	}
}
