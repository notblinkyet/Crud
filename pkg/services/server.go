package services

import (
	"net/http"
	"simple_crud/pkg/storage"
	"simple_crud/pkg/transport/handlers"
)

type server struct {
	addr      string
	r         *http.ServeMux
	dbStorage storage.Storage
}

func New(addr string, r *http.ServeMux, dbStorage storage.Storage) *server {
	return &server{addr: addr, r: r, dbStorage: dbStorage}
}

func (s *server) FillEndpoits() {
	s.r.Handle("/tasks/", handlers.RUDHandler(s.dbStorage))
	s.r.Handle("/tasks", handlers.RCHandler(s.dbStorage))
}

func (s *server) ListenAndServe() error {
	return http.ListenAndServe(s.addr, s.r)
}